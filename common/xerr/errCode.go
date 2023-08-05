package xerr

// 成功返回
const OK uint32 = 200

const InternalServerError uint32 = 10001
const ErrBind uint32 = 10002

const ErrValidation uint32 = 20001
const ErrDatabase uint32 = 20002
const ErrToken uint32 = 20003
const ErrGetQuery uint32 = 20004
const ErrGetParam uint32 = 20005
const ErrDelete uint32 = 20006

const ErrSecurityCheck uint32 = 20007

// Auth errors
const ErrAuthFailed uint32 = 20101
const ErrTokenInvalid uint32 = 20102

// user errors
const ErrCreateUser uint32 = 20201
const ErrUpdateUser uint32 = 20202
const ErrUserNotFound uint32 = 20203
const ErrGetUserInfo uint32 = 20204
const ErrUserInfo uint32 = 20205
const ErrAddSampleData uint32 = 20206
const ErrAddLicence uint32 = 20207

// comment errors
const ErrNotLiked uint32 = 20301
const ErrHasLiked uint32 = 20302
const ErrEvaluationList uint32 = 20303
const ErrCommentList uint32 = 20304
const ErrGetEvaluation uint32 = 20305
const ErrGetSubCommentInfo uint32 = 20306
const ErrGetParentCommentInfo uint32 = 20307
const ErrGetHotEvaluations uint32 = 20308
const ErrGetHistoryEvaluations uint32 = 20309
const ErrUpdateCourseInfo uint32 = 20310
const ErrDeleteComment uint32 = 20311
const ErrHasEvaluated uint32 = 20312
const ErrWordLimitation uint32 = 20313
const ErrGetEvaluationInfo uint32 = 20314

// table errors
const ErrTableExisting uint32 = 20401
const ErrClassExisting uint32 = 20402
const ErrGetTableInfo uint32 = 20403
const ErrGetClassInfo uint32 = 20404
const ErrCourseHasExisted uint32 = 20405
const ErrNewTable uint32 = 20406

// message errors
const ErrGetMessage uint32 = 20501

// search errors
const ErrSearchCourse uint32 = 20601

// upload errors
const ErrGetFile uint32 = 20701
const ErrUploadFile uint32 = 20702

// course errors
const ErrHistoryCourseExisting uint32 = 20801
const ErrGetSelfCourses uint32 = 20802
const ErrSavesDataToLocal uint32 = 20803
const ErrCreateHistoryCourse uint32 = 20804
const ErrFindUsingCourse uint32 = 20805
const ErrUsingCourseExisting uint32 = 20806
const ErrCreateReport uint32 = 20901

// collection errors
const ErrGetCollections uint32 = 21001
