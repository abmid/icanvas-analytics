import {getUser} from "@/helpers/utils"
import {login, logout} from "@/api/auth/auth"

const user = getUser()

const state = {
    currentUser : user,
    isLogin : !!user
}

// getters
const getters = {
    currentUser : (state) => {
        return state.currentUser
    },
    isLogin : (state) => {
        return state.isLogin
    }
}

// actions
// dispatch
const actions = {
    async login ({ commit }, creds) {
        let resLogin = await login(creds.email, creds.password)
        if(resLogin.status == 200){
            var payload = {
                email : resLogin.data.email,
                name : resLogin.data.name
            }
            commit("LOGIN_SUCCESS", payload)
        }

        return resLogin
    },
    async logout ({commit}){
        let resLogout = await logout()
        if (resLogout.status == 200) {
            localStorage.removeItem("icanvas_user")
            commit("LOGOUT")
        }
        return resLogout
    }
}

// mutations
// commit
const mutations = {
    LOGIN_SUCCESS : (state, payload) => {
        state.currentUser = Object.assign({}, payload)
        localStorage.setItem("icanvas_user", JSON.stringify(state.currentUser));
    },
    LOGOUT : (state) => {
        state.currentUser = null
        state.isLogin = false
    }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}