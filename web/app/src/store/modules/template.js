// initial state
import {getTemplateCurrent} from '@/helpers/utils';

const currentClass = getTemplateCurrent();

const state = {
    parrentClass : currentClass,
}

// getters
const getters = {
    parrentClass(state){
        return state.parrentClass;
    }
}

// actions
const actions = {}

// mutations
const mutations = {
    updateParrentClass(state, payload){
        localStorage.removeItem('template')
        state.parrentClass = payload;
        localStorage.setItem('template', payload);
    }

}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
