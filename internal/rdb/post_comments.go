// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package rdb

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// PostComment is an object representing the database table.
type PostComment struct {
	ID        uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    uint      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	PostID    uint      `boil:"post_id" json:"post_id" toml:"post_id" yaml:"post_id"`
	Content   string    `boil:"content" json:"content" toml:"content" yaml:"content"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *postCommentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L postCommentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PostCommentColumns = struct {
	ID        string
	UserID    string
	PostID    string
	Content   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	PostID:    "post_id",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

type whereHelperuint struct{ field string }

func (w whereHelperuint) EQ(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint) NEQ(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint) LT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint) LTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint) GT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint) GTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var PostCommentWhere = struct {
	ID        whereHelperuint
	UserID    whereHelperuint
	PostID    whereHelperuint
	Content   whereHelperstring
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperuint{field: "`post_comments`.`id`"},
	UserID:    whereHelperuint{field: "`post_comments`.`user_id`"},
	PostID:    whereHelperuint{field: "`post_comments`.`post_id`"},
	Content:   whereHelperstring{field: "`post_comments`.`content`"},
	CreatedAt: whereHelpernull_Time{field: "`post_comments`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`post_comments`.`updated_at`"},
}

// PostCommentRels is where relationship names are stored.
var PostCommentRels = struct {
	Post string
	User string
}{
	Post: "Post",
	User: "User",
}

// postCommentR is where relationships are stored.
type postCommentR struct {
	Post *Post
	User *User
}

// NewStruct creates a new relationship struct
func (*postCommentR) NewStruct() *postCommentR {
	return &postCommentR{}
}

// postCommentL is where Load methods for each relationship are stored.
type postCommentL struct{}

var (
	postCommentAllColumns            = []string{"id", "user_id", "post_id", "content", "created_at", "updated_at"}
	postCommentColumnsWithoutDefault = []string{"user_id", "post_id", "content"}
	postCommentColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	postCommentPrimaryKeyColumns     = []string{"id"}
)

type (
	// PostCommentSlice is an alias for a slice of pointers to PostComment.
	// This should generally be used opposed to []PostComment.
	PostCommentSlice []*PostComment
	// PostCommentHook is the signature for custom PostComment hook methods
	PostCommentHook func(context.Context, boil.ContextExecutor, *PostComment) error

	postCommentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	postCommentType                 = reflect.TypeOf(&PostComment{})
	postCommentMapping              = queries.MakeStructMapping(postCommentType)
	postCommentPrimaryKeyMapping, _ = queries.BindMapping(postCommentType, postCommentMapping, postCommentPrimaryKeyColumns)
	postCommentInsertCacheMut       sync.RWMutex
	postCommentInsertCache          = make(map[string]insertCache)
	postCommentUpdateCacheMut       sync.RWMutex
	postCommentUpdateCache          = make(map[string]updateCache)
	postCommentUpsertCacheMut       sync.RWMutex
	postCommentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var postCommentBeforeInsertHooks []PostCommentHook
var postCommentBeforeUpdateHooks []PostCommentHook
var postCommentBeforeDeleteHooks []PostCommentHook
var postCommentBeforeUpsertHooks []PostCommentHook

var postCommentAfterInsertHooks []PostCommentHook
var postCommentAfterSelectHooks []PostCommentHook
var postCommentAfterUpdateHooks []PostCommentHook
var postCommentAfterDeleteHooks []PostCommentHook
var postCommentAfterUpsertHooks []PostCommentHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PostComment) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postCommentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PostComment) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postCommentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PostComment) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postCommentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PostComment) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postCommentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PostComment) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postCommentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PostComment) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postCommentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PostComment) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postCommentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PostComment) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postCommentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PostComment) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postCommentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPostCommentHook registers your hook function for all future operations.
func AddPostCommentHook(hookPoint boil.HookPoint, postCommentHook PostCommentHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		postCommentBeforeInsertHooks = append(postCommentBeforeInsertHooks, postCommentHook)
	case boil.BeforeUpdateHook:
		postCommentBeforeUpdateHooks = append(postCommentBeforeUpdateHooks, postCommentHook)
	case boil.BeforeDeleteHook:
		postCommentBeforeDeleteHooks = append(postCommentBeforeDeleteHooks, postCommentHook)
	case boil.BeforeUpsertHook:
		postCommentBeforeUpsertHooks = append(postCommentBeforeUpsertHooks, postCommentHook)
	case boil.AfterInsertHook:
		postCommentAfterInsertHooks = append(postCommentAfterInsertHooks, postCommentHook)
	case boil.AfterSelectHook:
		postCommentAfterSelectHooks = append(postCommentAfterSelectHooks, postCommentHook)
	case boil.AfterUpdateHook:
		postCommentAfterUpdateHooks = append(postCommentAfterUpdateHooks, postCommentHook)
	case boil.AfterDeleteHook:
		postCommentAfterDeleteHooks = append(postCommentAfterDeleteHooks, postCommentHook)
	case boil.AfterUpsertHook:
		postCommentAfterUpsertHooks = append(postCommentAfterUpsertHooks, postCommentHook)
	}
}

// One returns a single postComment record from the query.
func (q postCommentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PostComment, error) {
	o := &PostComment{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rdb: failed to execute a one query for post_comments")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PostComment records from the query.
func (q postCommentQuery) All(ctx context.Context, exec boil.ContextExecutor) (PostCommentSlice, error) {
	var o []*PostComment

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "rdb: failed to assign all query results to PostComment slice")
	}

	if len(postCommentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PostComment records in the query.
func (q postCommentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to count post_comments rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q postCommentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "rdb: failed to check if post_comments exists")
	}

	return count > 0, nil
}

// Post pointed to by the foreign key.
func (o *PostComment) Post(mods ...qm.QueryMod) postQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.PostID),
	}

	queryMods = append(queryMods, mods...)

	query := Posts(queryMods...)
	queries.SetFrom(query.Query, "`posts`")

	return query
}

// User pointed to by the foreign key.
func (o *PostComment) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// LoadPost allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (postCommentL) LoadPost(ctx context.Context, e boil.ContextExecutor, singular bool, maybePostComment interface{}, mods queries.Applicator) error {
	var slice []*PostComment
	var object *PostComment

	if singular {
		object = maybePostComment.(*PostComment)
	} else {
		slice = *maybePostComment.(*[]*PostComment)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &postCommentR{}
		}
		args = append(args, object.PostID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &postCommentR{}
			}

			for _, a := range args {
				if a == obj.PostID {
					continue Outer
				}
			}

			args = append(args, obj.PostID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`posts`), qm.WhereIn(`posts.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Post")
	}

	var resultSlice []*Post
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Post")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for posts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for posts")
	}

	if len(postCommentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Post = foreign
		if foreign.R == nil {
			foreign.R = &postR{}
		}
		foreign.R.PostComments = append(foreign.R.PostComments, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PostID == foreign.ID {
				local.R.Post = foreign
				if foreign.R == nil {
					foreign.R = &postR{}
				}
				foreign.R.PostComments = append(foreign.R.PostComments, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (postCommentL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybePostComment interface{}, mods queries.Applicator) error {
	var slice []*PostComment
	var object *PostComment

	if singular {
		object = maybePostComment.(*PostComment)
	} else {
		slice = *maybePostComment.(*[]*PostComment)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &postCommentR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &postCommentR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`users.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(postCommentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.PostComments = append(foreign.R.PostComments, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.PostComments = append(foreign.R.PostComments, local)
				break
			}
		}
	}

	return nil
}

// SetPost of the postComment to the related item.
// Sets o.R.Post to related.
// Adds o to related.R.PostComments.
func (o *PostComment) SetPost(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Post) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `post_comments` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"post_id"}),
		strmangle.WhereClause("`", "`", 0, postCommentPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PostID = related.ID
	if o.R == nil {
		o.R = &postCommentR{
			Post: related,
		}
	} else {
		o.R.Post = related
	}

	if related.R == nil {
		related.R = &postR{
			PostComments: PostCommentSlice{o},
		}
	} else {
		related.R.PostComments = append(related.R.PostComments, o)
	}

	return nil
}

// SetUser of the postComment to the related item.
// Sets o.R.User to related.
// Adds o to related.R.PostComments.
func (o *PostComment) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `post_comments` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, postCommentPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &postCommentR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			PostComments: PostCommentSlice{o},
		}
	} else {
		related.R.PostComments = append(related.R.PostComments, o)
	}

	return nil
}

// PostComments retrieves all the records using an executor.
func PostComments(mods ...qm.QueryMod) postCommentQuery {
	mods = append(mods, qm.From("`post_comments`"))
	return postCommentQuery{NewQuery(mods...)}
}

// FindPostComment retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPostComment(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*PostComment, error) {
	postCommentObj := &PostComment{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `post_comments` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, postCommentObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rdb: unable to select from post_comments")
	}

	return postCommentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PostComment) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("rdb: no post_comments provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(postCommentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	postCommentInsertCacheMut.RLock()
	cache, cached := postCommentInsertCache[key]
	postCommentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			postCommentAllColumns,
			postCommentColumnsWithDefault,
			postCommentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(postCommentType, postCommentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(postCommentType, postCommentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `post_comments` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `post_comments` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `post_comments` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, postCommentPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "rdb: unable to insert into post_comments")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == postCommentMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "rdb: unable to populate default values for post_comments")
	}

CacheNoHooks:
	if !cached {
		postCommentInsertCacheMut.Lock()
		postCommentInsertCache[key] = cache
		postCommentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PostComment.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PostComment) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	postCommentUpdateCacheMut.RLock()
	cache, cached := postCommentUpdateCache[key]
	postCommentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			postCommentAllColumns,
			postCommentPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("rdb: unable to update post_comments, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `post_comments` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, postCommentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(postCommentType, postCommentMapping, append(wl, postCommentPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to update post_comments row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to get rows affected by update for post_comments")
	}

	if !cached {
		postCommentUpdateCacheMut.Lock()
		postCommentUpdateCache[key] = cache
		postCommentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q postCommentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to update all for post_comments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to retrieve rows affected for post_comments")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PostCommentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("rdb: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postCommentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `post_comments` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, postCommentPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to update all in postComment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to retrieve rows affected all in update all postComment")
	}
	return rowsAff, nil
}

var mySQLPostCommentUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PostComment) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("rdb: no post_comments provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(postCommentColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPostCommentUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	postCommentUpsertCacheMut.RLock()
	cache, cached := postCommentUpsertCache[key]
	postCommentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			postCommentAllColumns,
			postCommentColumnsWithDefault,
			postCommentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			postCommentAllColumns,
			postCommentPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("rdb: unable to upsert post_comments, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "post_comments", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `post_comments` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(postCommentType, postCommentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(postCommentType, postCommentMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "rdb: unable to upsert for post_comments")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == postCommentMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(postCommentType, postCommentMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "rdb: unable to retrieve unique values for post_comments")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "rdb: unable to populate default values for post_comments")
	}

CacheNoHooks:
	if !cached {
		postCommentUpsertCacheMut.Lock()
		postCommentUpsertCache[key] = cache
		postCommentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PostComment record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PostComment) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("rdb: no PostComment provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), postCommentPrimaryKeyMapping)
	sql := "DELETE FROM `post_comments` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to delete from post_comments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to get rows affected by delete for post_comments")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q postCommentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("rdb: no postCommentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to delete all from post_comments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to get rows affected by deleteall for post_comments")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PostCommentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(postCommentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postCommentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `post_comments` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, postCommentPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to delete all from postComment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to get rows affected by deleteall for post_comments")
	}

	if len(postCommentAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PostComment) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPostComment(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PostCommentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PostCommentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postCommentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `post_comments`.* FROM `post_comments` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, postCommentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "rdb: unable to reload all in PostCommentSlice")
	}

	*o = slice

	return nil
}

// PostCommentExists checks if the PostComment row exists.
func PostCommentExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `post_comments` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "rdb: unable to check if post_comments exists")
	}

	return exists, nil
}
