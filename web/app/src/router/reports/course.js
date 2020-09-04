const DashboardReportCourse = () => import("@/views/dashboard/reports/Course.vue")

export default [
    {
        path: '/report/course',
        name: 'dashboard.report.course',
        component: DashboardReportCourse,
        meta : {
          requiredAuth : true,
          title : "Course Reports"
        }
    },
]