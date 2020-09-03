import request from '@/helpers/request';

export function courseAnalytics(filter) {
    var accountID = filter.account_id == "all" ? null : filter.account_id
    var analyticsTeacher = filter.analytics_teacher
    var date = filter.date == "" ? null : filter.date
    var page = filter.page == "" ? 1 : filter.page
    var limit = filter.limit == "" ? 10 : filter.limit
    var orderBy = filter.order_by == "" ? "desc" : filter.order_by
    var params = {
        account_id : accountID,
        analytics_teacher : analyticsTeacher,
        date : date,
        page : page,
        limit : limit,
        order_by : orderBy

    }
    return request({
        url: `/analytics/courses`,
        params: params,
        method: "GET",
        withCredentials: true
    })
}