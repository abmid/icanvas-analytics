
const DashboardSetting = () => import("@/views/dashboard/settings/Canvas")

export default [
    {
        path: '/setting',
        name: 'dashboard.setting',
        component: DashboardSetting,
        meta : {
          requiredAuth: true,
          title: "Setting"
        }
      } 
]