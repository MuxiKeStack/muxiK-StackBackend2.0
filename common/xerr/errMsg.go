package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[InternalServerError] = "Internal server error"
	message[ErrBind] = "Error occurred while binding the request body to the struct."

	message[ErrValidation] = "Validation failed."
	message[ErrDatabase] = "Database error."
	message[ErrToken] = "Error occurred while signing the JSON web token."
	message[ErrGetQuery] = "Error occurred while getting query. "
	message[ErrGetParam] = "Error occurred while getting path params. "
	message[ErrDelete] = "Error occurred while deleting sth. "

	message[ErrSecurityCheck] = "Error occurred in security check"

	// Auth errors
	message[ErrAuthFailed] = "The sid or password was incorrect."
	message[ErrTokenInvalid] = "The token was invalid."

	// user errors
	message[ErrCreateUser] = "Error occurred in creating user."
	message[ErrUpdateUser] = "Error occurred in updating user"
	message[ErrUserNotFound] = "The user was not found."
	message[ErrGetUserInfo] = "Error in getting user info"
	message[ErrUserInfo] = "The user information json cannot be null"
	message[ErrAddSampleData] = "Error occurred in joining project"
	message[ErrAddLicence] = "Error occurred in grade licence getting"

	// comment errors
	message[ErrNotLiked] = "User has not liked yet. "
	message[ErrHasLiked] = "User has already liked. "
	message[ErrEvaluationList] = "Error occurred while getting evaluation list. "
	message[ErrCommentList] = "Error occurred while getting comment list. "
	message[ErrGetEvaluation] = "Error occurred while getting evaluation."
	message[ErrGetSubCommentInfo] = "Error occurred while getting subComment info"
	message[ErrGetParentCommentInfo] = "Error occurred while getting parent comment info"
	message[ErrGetHotEvaluations] = "Error occurred while getting hot evaluations"
	message[ErrGetHistoryEvaluations] = "Error occurred while getting history evaluations"
	message[ErrUpdateCourseInfo] = "Error occurred while updating course's info"
	message[ErrDeleteComment] = "Error occurred while deleting a comment "
	message[ErrHasEvaluated] = "User has evaluated the course"
	message[ErrWordLimitation] = "Word limit exceeded"
	message[ErrGetEvaluationInfo] = "Error occurred while getting evaluation info"

	// table errors
	message[ErrTableExisting] = "The table does not exist"
	message[ErrClassExisting] = "The class does not exist"
	message[ErrGetTableInfo] = "Error occurred in getting table info. "
	message[ErrGetClassInfo] = "Error occurred in getting class info."
	message[ErrCourseHasExisted] = "This course has already existed in the table."
	message[ErrNewTable] = "Error occurred while creating a new table "

	// message errors
	message[ErrGetMessage] = "Error occurred in getting message list"

	// search errors
	message[ErrSearchCourse] = "Error occured in searching courses."

	// upload errors
	message[ErrGetFile] = "Error occurred in getting file from FormFile()"
	message[ErrUploadFile] = "Error occurred in uploading file to oss"

	// course errors
	message[ErrHistoryCourseExisting] = "History Course does not exist."
	message[ErrGetSelfCourses] = "Error occurred in getting self courses"
	message[ErrSavesDataToLocal] = "Error occurred in Saving data to local"
	message[ErrCreateHistoryCourse] = "Error occurred in creating a new history course"
	message[ErrFindUsingCourse] = "Error occurred in finding the specific using course"
	message[ErrUsingCourseExisting] = "Using Course does not exist."

	// report errors
	message[ErrCreateReport] = "Error occurred in creating new report."

	// collection errors
	message[ErrGetCollections] = "Error occurred in getting collections"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
