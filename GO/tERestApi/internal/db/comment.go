package db

import (
	"context"
	"database/sql"
	"edgerestapi/internal/comment"
	"fmt"

	"github.com/gofrs/uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:   c.ID,
		Slug: c.Slug.String,
		Body: c.Body.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(ctx, `SELECT id, slug, body, author FROM comments WHERE id = $1`, uuid)
	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error feetching the comment by uuid: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Slug, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments
		(id, slug, author, body) 
		VALUES 
		(:id, :slug, :author, :body)`,
		postRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error posting the comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("error closing the rows: %w", err)
	}

	return cmt, nil
}

func (d *Database) DeleteComment(ctx context.Context, cmt comment.Comment) error {
	_, err := d.Client.ExecContext(ctx, `DELETE FROM comments WHERE id = $1`, cmt.ID)
	if err != nil {
		return fmt.Errorf("error deleting the comment: %w", err)
	}

	return nil
}

func (d *Database) UpdateComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmtRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Slug, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		slug = :slug,
		author = :author,
		body = :body
		WHERE id = :id`,
		cmtRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error updating the comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("error closing the rows: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}
