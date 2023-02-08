import request from '@/utils/request';

// 获取数据
export function panoramaImgList(query) {
  return request({
    url: '/system/panoramaImg/List',
    method: 'get',
    params: query,
  });
}
// 添加图片
export function panoramaImgAdd(img) {
  return request({
    url: '/system/panoramaImg/Add',
    method: 'post',
    data: {
      picture: img,
    },
  });
}
// 删除图片
export function panoramaImgDelete(id) {
  return request({
    url: '/system/panoramaImg/Delete',
    method: 'delete',
    data: { ids: id },
  });
}
