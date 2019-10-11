export default {
  namespace: 'file',
  state: [],

  effects: {
    *query({ payload, callback }, { call, put }) {
      const response = yield call("", payload);
      yield put({
        type: 'save',
        payload: {
          response: response,
        },
      });
      if (callback) {
        callback(response);
      }
    },
    *delete({ payload, callback }, { call, put }) {
      const response = yield call("", payload);
      yield put({
        type: 'save',
        payload: {
          response: response,
        },
      });
      if (callback) {
        callback(response);
      }
    },
  },
  reducers: {
    save(state, { payload }) {
      return {
        ...state,
        ...payload,
      };
    },
  },
};
