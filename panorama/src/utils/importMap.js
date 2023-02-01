const modules = import.meta.globEager('../../node_modules/echarts/map/js/province/*.js');
export function filterMap(name) {
  let a = modules[`../../node_modules/echarts/map/js/province/${name}.js`];
  return a;
}
