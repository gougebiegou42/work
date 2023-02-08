<template>
  <div id="map" @dblclick="back()"></div>
</template>

<script>
// 地区映射对象
import { provinces } from '@/utils/provincesmap';
import { cityMap } from '@/utils/citymap';
export default {
  data() {
    return {
      chart: null,
      showMap: false,
      // 经度
      malng: 0,
      // 纬度
      maplat: 0,
      //直辖市和特别行政区-只有二级地图，没有三级地图
      special: ["北京", "天津", "上海", "重庆", "香港", "澳门"],
      mapData: [],
      //初始化绘制全国地图配置
      option: {
        // backgroundColor: "#000",
        tooltip: {
          trigger: "item",
          formatter: "{b}",
        },
        animationDuration: 1000,
        animationEasing: "cubicOut",
        animationDurationUpdate: 1000,
      }
    }
  },
  mounted() {
    let that = this;
    // this.$store.commit("region/hideDisable", true);
    // console.log(this.$store.state.region.aroundHide);


    // this.initJson();
    //地图容器
    this.chart = that.$echarts.init(document.getElementById("map"));
    // 初始化中国地图
    this.loadChina();
    // 注册点击事件
    this.chart.on('click', async (params) => {
      // 判断点击的是不是省份
      if (params.name in provinces) {
        // 如果点击的是34个省、市、自治区，绘制选中地区的二级地图
        let data = await import(`../assets/map/province/${provinces[params.name]}.json`);
        console.log(data.default);
        if (that.special.indexOf(params.name) >= 0) {
          this.$store.commit("region/hideDisable", false);
          this.$store.commit('region/selectDisable', params.name);
        }
        // 注册省份地图数据
        that.$echarts.registerMap(params.name, data.default);
        // 添加省份对应的城市数据
        let d = [];
        for (let i = 0; i < data.default.features.length; i++) {
          d.push({
            name: data.default.features[i].properties.name
          })
        }
        // 重新绘制地图
        that.renderMap(params.name, d);
      } else if (params.seriesName in provinces) {
        // 判断点击的是不是城市
        // 如果是【直辖市/特别行政区】只有二级下钻
        if (that.special.indexOf(params.seriesName) >= 0) {
          console.log(params.seriesName);
          // 打开百度地图 
          // let { data: { result: { location: { lat, lng } } } } = await axios.get('/baidu/geocoding/v3/', {
          //   params: {
          //     address: params.seriesName + params.name,
          //     ak: 'CtDFZnPIhj7MjWLllGtFNyLpT4MpgGwv',
          //     output: 'json'
          //   }
          // })
          // maplng.value = lng
          // maplat.value = lat
          // showMap.value = true
        } else {
          //显示县级地图
          // let { data } = await axios.get('/map/city/' + cityMap[params.name] + '.json');
          let data = await import(`../assets/map/city/${cityMap[params.name]}.json`)
          console.log(params.name)
          this.$store.commit("region/hideDisable", false);
          this.$store.commit('region/selectDisable', params.name);
          // 注册城市地图数据
          that.$echarts.registerMap(params.name, data.default);
          // 添加城市对应的区县数据
          let d = [];
          for (let i = 0; i < data.default.features.length; i++) {
            d.push({
              name: data.default.features[i].properties.name
            })
          }
          // 重新绘制地图
          that.renderMap(params.name, d);
        }
      }
      // }else{
      //   // 打开百度地图 
      //   let {data:{result:{location:{lat,lng}}}} = await axios.get('/baidu/geocoding/v3/',{
      //     params:{
      //       address:params.seriesName+params.name,
      //       ak:'sICFGuAAPisdVnB5Bx6ILAADGPnRDusH',
      //       output:'json'
      //     }
      //   })
      //   maplng.value = lng
      //   maplat.value = lat
      //   showMap.value = true
      // }
    })
    window.addEventListener(
      'resize',
      function() {
        that.chart.resize();
      }
    );
  },
  methods: {
    async initJson() {
      let a = 'guangdong';
      // let json = import.meta.glob('../assets/map/province/*.json');
      // console.log(json);

      let module = await import(`../assets/map/province/${a}.json`);
      console.log(module.default);
      // for (const key in json) {
      // const result = key.match(/.*\/(.+).json$/);

      // console.log(result);
      // if (result) {

      // }
      // if (a in provinces) {
      // json[key]().then(res => {
      //   console.log(res.default);
      // })
      // }
      // }
    },
    //绘制地图的方法
    renderMap(map, data) {
      let that = this
      this.option.series = [
        {
          name: map,
          type: "map",
          mapType: map,
          roam: false,
          nameMap: {
            china: "中国",
          },
          label: {
            normal: {
              show: true,
              textStyle: {
                color: "#d2ddea",
                fontSize: 13,
              },
            },
            emphasis: {
              show: true,
              textStyle: {
                color: "#fff",
                fontSize: 13,
              },
            },
          },
          itemStyle: {
            normal: {
              areaColor: "#003181",
              borderColor: "dodgerblue",
            },
            emphasis: {
              areaColor: "#3f74b9",
            },
          },
          data: data,
        },
      ];
      //渲染地图
      this.chart.setOption(this.option);
    },
    // 初始化中国地图
    async loadChina() {
      // let { data } = await axios.get("/map/china.json");
      const module = await import('../assets/map/china.json');
      // let module = await import(`../assets/map/province/${provinces['浙江']}.json`);
      // let module = await import(`../assets/map/city/${cityMap['广州市']}.json`)
      let d = [];
      for (let i = 0; i < module.default.features.length; i++) {
        d.push({
          name: module.default.features[i].properties.name,
        });
      }
      this.mapData = d;
      //注册地图
      this.$echarts.registerMap("china", module.default);
      // this.$echarts.registerMap("广州市", module.default);
      //绘制地图
      this.renderMap('china', this.mapData);
      // this.renderMap('广州市', this.mapData);
    },
    // 选择调入对应区域
    async selectCity(region, type) {
      let data = null;
      if (type === 1) {
        data = await import(`../assets/map/province/${provinces[region]}.json`)
      } else if (type === 2) {
        data = await import(`../assets/map/city/${cityMap[region]}.json`)
      }
      // 注册城市地图数据
      this.$echarts.registerMap(region, data.default);
      // 添加城市对应的区县数据
      let d = [];
      for (let i = 0; i < data.default.features.length; i++) {
        d.push({
          name: data.default.features[i].properties.name
        })
      }
      // 重新绘制地图
      this.renderMap(region, d);
    },
    // 返回
    back() {
      this.renderMap('china', this.mapData)
      this.$store.commit("region/hideDisable", true);
      this.$store.commit("region/selectDisable", 'china');
    }

  },
  watch: {
    '$store.state.region.selectName'(val) {
      if (val === 'china') {
        this.renderMap('china', this.mapData)
      } else {
        let shearStr = val.slice(0, 2);
        if (this.special.indexOf(shearStr) >= 0) {
          this.selectCity(shearStr, 1)
        } else {
          this.selectCity(val, 2)
        }
      }
    }
  }
}
</script>

<style lang="less" scoped>
#map {
  width: 100%;
  height: 100%;
}
</style>