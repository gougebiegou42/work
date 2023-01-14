import { get } from '@/utils/request';

export default class User {
  /**
   * 登录
   * @param {String} query 用户名
   * @returns
   */
  static async listSylas(query) {
    return get('/system/sylas/List', {
      query,
    });
  }
}
