const state = {
  aroundHide: true,
  selectName: '',
};

const mutations = {
  hideDisable(state, val) {
    state.aroundHide = val;
  },
  selectDisable(state, val) {
    state.selectName = val;
  },
};

const actions = {
  hideDisable({ commit }) {
    commit('increment');
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
};
