// Code generated by goctl. DO NOT EDIT.

package cmsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	cmsSubjectFieldNames          = builder.RawFieldNames(&CmsSubject{})
	cmsSubjectRows                = strings.Join(cmsSubjectFieldNames, ",")
	cmsSubjectRowsExpectAutoSet   = strings.Join(stringx.Remove(cmsSubjectFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	cmsSubjectRowsWithPlaceHolder = strings.Join(stringx.Remove(cmsSubjectFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	cmsSubjectModel interface {
		Insert(ctx context.Context, data *CmsSubject) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*CmsSubject, error)
		Update(ctx context.Context, data *CmsSubject) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCmsSubjectModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CmsSubject struct {
		Id              int64          `db:"id"`
		CategoryId      int64          `db:"category_id"`
		Title           string         `db:"title"`
		Pic             string         `db:"pic"`           // 专题主图
		ProductCount    int64          `db:"product_count"` // 关联产品数量
		RecommendStatus int64          `db:"recommend_status"`
		CreateTime      time.Time      `db:"create_time"`
		CollectCount    int64          `db:"collect_count"`
		ReadCount       int64          `db:"read_count"`
		CommentCount    int64          `db:"comment_count"`
		AlbumPics       string         `db:"album_pics"` // 画册图片用逗号分割
		Description     sql.NullString `db:"description"`
		ShowStatus      int64          `db:"show_status"` // 显示状态：0->不显示；1->显示
		Content         string         `db:"content"`
		ForwardCount    int64          `db:"forward_count"` // 转发数
		CategoryName    string         `db:"category_name"` // 专题分类名称
	}
)

func newCmsSubjectModel(conn sqlx.SqlConn) *defaultCmsSubjectModel {
	return &defaultCmsSubjectModel{
		conn:  conn,
		table: "`cms_subject`",
	}
}

func (m *defaultCmsSubjectModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCmsSubjectModel) FindOne(ctx context.Context, id int64) (*CmsSubject, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cmsSubjectRows, m.table)
	var resp CmsSubject
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCmsSubjectModel) Insert(ctx context.Context, data *CmsSubject) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, cmsSubjectRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CategoryId, data.Title, data.Pic, data.ProductCount, data.RecommendStatus, data.CollectCount, data.ReadCount, data.CommentCount, data.AlbumPics, data.Description, data.ShowStatus, data.Content, data.ForwardCount, data.CategoryName)
	return ret, err
}

func (m *defaultCmsSubjectModel) Update(ctx context.Context, data *CmsSubject) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, cmsSubjectRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.CategoryId, data.Title, data.Pic, data.ProductCount, data.RecommendStatus, data.CollectCount, data.ReadCount, data.CommentCount, data.AlbumPics, data.Description, data.ShowStatus, data.Content, data.ForwardCount, data.CategoryName, data.Id)
	return err
}

func (m *defaultCmsSubjectModel) tableName() string {
	return m.table
}