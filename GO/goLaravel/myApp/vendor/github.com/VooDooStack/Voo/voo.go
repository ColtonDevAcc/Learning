package voo

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/VooDooStack/Voo/render"
	"github.com/VooDooStack/Voo/session"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "0.0.1"

type Voo struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	JetViews *jet.Set
	Session *scs.SessionManager
	config   config
}

type config struct {
	port        string
	renderer    string
	cookie      cookieConfig
	sessionType string
}

func (v *Voo) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"migrations", "handlers", "logs", "data", "tmp", "views", "public", "tmp", "middlewares"},
	}
	err := v.Init(pathConfig)
	if err != nil {

		return err
	}

	err = v.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	//! //=========== read .env file ===========//
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	//! //=========== create loggers ===========//
	infoLog, errorLog := v.startLoggers()
	v.InfoLog = infoLog
	v.ErrorLog = errorLog
	v.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	v.Version = version
	v.RootPath = rootPath
	v.Routes = v.routes().(*chi.Mux)

	v.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
		cookie: cookieConfig{
			name:     os.Getenv("COOKIE_NAME"),
			lifetime: os.Getenv("COOKIE_LIFETIME"),
			persist:  os.Getenv("COOKIE_PERSIST"),
			secure:   os.Getenv("COOKIE_SECURE"),
			domain:   os.Getenv("COOKIE_DOMAIN"),
		},
		sessionType: os.Getenv("SESSION_TYPE"),
	}

	//! //=========== create the session ===========//
	sess := session.Session{
		CookieLifetime: v.config.cookie.lifetime,
		CookieName:     v.config.cookie.name,
		CookiePersist:  v.config.cookie.persist,
		SessionType:    v.config.sessionType,
		CookieDomain:   v.config.cookie.domain,
	}

	v.Session = sess.InitSession()

	//! //=========== create renderer ===========//
	var views = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	v.JetViews = views

	v.createRenderer()

	return nil
}

func (v *Voo) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		//? if folder does not exits, create it
		err := v.CreateDirIfNotExits(root + "/" + path)
		if err != nil {
			return err
		}
	}

	return nil
}

//! //=========== listen and serve starts server ===========//
func (v *Voo) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     v.ErrorLog,
		Handler:      v.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	v.InfoLog.Printf("listening on port %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	if err != nil {
		v.ErrorLog.Fatal(err)
	}
}

func (v *Voo) checkDotEnv(path string) error {
	err := v.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

func (v *Voo) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t ", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (v *Voo) createRenderer() {
	myRenderer := render.Render{
		Renderer: v.config.renderer,
		RootPath: v.RootPath,
		Port:     v.config.port,
		JetViews: v.JetViews,
	}
	v.Render = &myRenderer
}
