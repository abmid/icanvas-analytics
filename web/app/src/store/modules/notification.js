/*
 * File Created: Tuesday, 20th October 2020 1:58:41 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

const state = {
    notification : []
}

// getters
const getters = {
    notification : (state) => {
        return state.notification
    }
}

// actions
// dispatch
const actions = {
    addNotification({commit}, payload){
        commit("ADD_NOTIFICATION", payload)
    }
}

// mutations
// commit
const mutations = {
    ADD_NOTIFICATION : (state, payload) => {
        state.notification.unshift(payload)
    },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}