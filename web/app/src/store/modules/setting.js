import {isExistsCanvasConfig} from "@/api/settings/setting"

const state = {
    canvasConfig : null
}

// getters
const getters = {
    canvasConfig : (state) => {
        return state.canvasConfig
    }
}

// actions
// dispatch
const actions = {
    async isExistsCanvasConfig({commit}){
        let res = await isExistsCanvasConfig()
        if (res.status == 200) {
            commit("CANVAS_CONFIG_EXISTS", res.data)
        }
        return res
    }
}

// mutations
// commit
const mutations = {
    CANVAS_CONFIG_EXISTS : (state, payload) => {
        state.canvasConfig = {
            token : payload.token,
            url : payload.url
        }
    },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}