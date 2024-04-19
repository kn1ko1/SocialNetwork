package comments

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes comments related to CommentId from the COMMENTS table
func DeleteCommentById(db *sql.DB, commentId int) error {
	query := `DELETE FROM "COMMENTS" WHERE "CommentId" = (?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing delete comment statement.", err)
		return err
	}
	_, err = stmt.Exec(commentId)
	if err != nil {
		utils.HandleError("Error executing delete comment statement.", err)
		return err
	}
	return nil
}
