<template>
  <div class="panorama">
    <!-- 头部栏 -->
    <div class="header">
      <div class="top">
        <div class="top_left">
          <div class="back"></div>
          <div class="content">分散式\全国</div>
        </div>
        <h1>住房租赁数字化经营全景图</h1>
        <div class="top_right">
          <div class="date">{{ currentDate }}</div>
          <div class="time">{{ currentTime }}</div>
        </div>
      </div>
      <!-- <div class="select">
        全国
        <div class="select_icon"></div>
      </div> -->
      <SelectRegion></SelectRegion>
    </div>
    <!-- 主体内容 -->
    <div class="main">
      <div class="child main_left " :class="{ 'animate__fadeInLeft': aroundHide }">
        <!-- 累计项目数 -->
        <div class="total_data" v-show="aroundHide">
          <div class="total_left">
            <div class="project_icon"></div>
            <div class="total_progress"></div>
          </div>
          <div class="total_right">
            <div class="r_top">{{ pcount }}</div>
            <div class="r_bottom">累计项目数</div>
          </div>
          <div class="chart"></div>
        </div>
        <!-- 房源项目规模 -->
        <div class="project_scale" v-show="aroundHide">
          <div class="p_header">
            <div class="p_icon"></div>
            <div class="p_title">房源项目规模</div>
          </div>
          <div class="p_main">
            <div class="row pm_1">
              <div class="row_info">
                <div class="info_left">
                  <div class="info_icon icon1"></div>
                  <div class="info_title">累计房源数</div>
                </div>
                <div class="info_right">{{ housesCountCal }} 万间</div>
              </div>
              <div class="row_progress progress1">
                <div style="width: 100%; height: 100%; overflow: hidden;">
                  <img :style="{ left: `calc(-${rowLeft1}%)` }" src="../assets/images/1房源项目规模/黄.png">
                </div>
              </div>
            </div>
            <div class="row pm_2">
              <div class="row_info">
                <div class="info_left">
                  <div class="info_icon icon2"></div>
                  <div class="info_title">当前持有房源数</div>
                </div>
                <div class="info_right">{{ currentHousesCal }} 万间</div>
              </div>
              <div class="row_progress progress2">
                <div style="width: 100%; height: 100%; overflow: hidden;">
                  <img :style="{ left: `calc(-${rowLeft2}%)` }" src="../assets/images/1房源项目规模/青.png">
                </div>
              </div>
            </div>
            <div class="row pm_3">
              <div class="row_info">
                <div class="info_left">
                  <div class="info_icon icon3"></div>
                  <div class="info_title">可运营房源数</div>
                </div>
                <div class="info_right">{{ operableHousesCal }} 万间</div>
              </div>
              <div class="row_progress progress3">
                <div style="width: 100%; height: 100%; overflow: hidden;">
                  <img :style="{ left: `calc(-${rowLeft3}%)` }" src="../assets/images/1房源项目规模/绿.png">
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="child main_middle">
        <!-- 地图 -->
        <!-- <div class="map_chart"></div> -->
        <Map></Map>
      </div>
      <div class="child main_right" :class="{ 'animate__fadeInRight': aroundHide }">
        <!-- 运营效率 -->
        <div class="run_efficiency" v-show="aroundHide">
          <div class="r_header">
            <img src="../assets/images/2运营效率/运营.png" alt="">
            <div class="r_title">运营效率</div>
          </div>
          <div class="r_main">
            <div class="rm_ripple">
              <div class="ripple">
                <div class="ripple_bar">
                  <div class="bar_box">
                    <div class="bar_inner" :style="barProgress()">
                      <svg xmlns="http://www.w3.org/2000/svg" version="1.0" viewBox="0 0 600 140" class="box_waves">
                        <path d="M 0 70 Q 75 20,150 70 T 300 70 T 450 70 T 600 70 L 600 140 L 0 140 L 0 70Z">
                        </path>
                      </svg>
                      <svg xmlns="http://www.w3.org/2000/svg" version="1.0" viewBox="0 0 600 140" class="box_waves">
                        <path d="M 0 70 Q 75 20,150 70 T 300 70 T 450 70 T 600 70 L 600 140 L 0 140 L 0 70Z">
                        </path>
                      </svg>
                      <svg xmlns="http://www.w3.org/2000/svg" version="1.0" viewBox="0 0 600 140" class="box_waves">
                        <path d="M 0 70 Q 75 20,150 70 T 300 70 T 450 70 T 600 70 L 600 140 L 0 140 L 0 70Z">
                        </path>
                      </svg>
                    </div>
                  </div>
                  <div class="ripple_num">{{ rentalRate }} <div>%</div>
                  </div>
                </div>
                <div class="ripple_tag">出租率</div>
              </div>
              <div class="ripple">
                <div class="ripple_bar">
                  <div class="bar_box">
                    <div class="bar_inner" :style="barProgress()">
                      <svg xmlns="http://www.w3.org/2000/svg" version="1.0" viewBox="0 0 600 140" class="box_waves">
                        <path d="M 0 70 Q 75 20,150 70 T 300 70 T 450 70 T 600 70 L 600 140 L 0 140 L 0 70Z">
                        </path>
                      </svg>
                      <svg xmlns="http://www.w3.org/2000/svg" version="1.0" viewBox="0 0 600 140" class="box_waves">
                        <path d="M 0 70 Q 75 20,150 70 T 300 70 T 450 70 T 600 70 L 600 140 L 0 140 L 0 70Z">
                        </path>
                      </svg>
                      <svg xmlns="http://www.w3.org/2000/svg" version="1.0" viewBox="0 0 600 140" class="box_waves">
                        <path d="M 0 70 Q 75 20,150 70 T 300 70 T 450 70 T 600 70 L 600 140 L 0 140 L 0 70Z">
                        </path>
                      </svg>
                    </div>
                  </div>
                  <div class="ripple_num">{{ rentalPrice }} <div>元/间/月</div>
                  </div>
                </div>
                <div class="ripple_tag">出租均价</div>
              </div>
            </div>
            <div class="rm_group">
              <div class="group group_left">
                <div class="group_info">
                  <div class="gi_top">{{ serviceEnt }}<span>家</span></div>
                  <div class="gi_bottom">累计服务企业数</div>
                </div>
                <img src="../assets/images/2运营效率/企业.png" alt="">
              </div>
              <div class="group group_right">
                <img src="../assets/images/2运营效率/租客.png" alt="">
                <div class="group_info">
                  <div class="gi_top">{{ serviceTenantCal }}<span>万人</span></div>
                  <div class="gi_bottom">累计服务租客数</div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <!-- 产业生态 -->
        <div class="product_ecology" v-show="aroundHide">
          <div class="pe_header">
            <img src="../assets/images/3产业生态/产业生态.png" alt="">
            <div class="pe_title">产品生态</div>
          </div>
          <div class="pe_main">
            <div class="pe_info pe_left">
              <div class="pi_top">{{ smartDevicesCal }}<span>万个</span></div>
              <div class="pi_bottom">接入智能设备数</div>
            </div>
            <div class="pe_info pe_right">
              <div class="pi_top">{{ cooperativeOp }}<span>家</span></div>
              <div class="pi_bottom">合作运营商</div>
            </div>
            <div class="pe_info pe_bottom">
              <div class="pi_top">{{ allianceEnt }}<span>家</span></div>
              <ScrollNum :value="'6952'"></ScrollNum>
              <div class="pi_bottom">入库联盟企业数</div>
            </div>
            <div class="pe_chart">
              <div class="pe_mask"></div>
              <div class="pe_chart_1">
                <div class="inner_1"></div>
              </div>
              <div class="pe_chart_2">
                <div class="inner_2"></div>
              </div>
              <div class="pe_chart_3">
                <div class="inner_3"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Prefecture v-show="!aroundHide"></Prefecture>
    </div>
    <!-- 底部切换 -->
    <div class="footer">
      <div class="footer_focus footer_left">
        集中式
      </div>
      <div class="footer_focus footer_right">
        分散式
      </div>
    </div>
  </div>
</template>

<script>
import init from '../utils/initChart';
import '@/utils/china.js';
import Prefecture from '@/components/Prefecture.vue'
import SelectRegion from '@/components/SelectRegion.vue'
import Map from '@/components/map.vue'
import ScrollNum from '@/components/scrollNum.vue'
import { panoramaList } from '@/api/list';


export default {
  components: {
    Prefecture,
    SelectRegion,
    Map,
    ScrollNum
  },
  data() {
    return {
      // bar进度条垂直位移
      barBottom: 50,
      // row进度条水平位移
      rowLeft1: 100,
      rowLeft2: 100,
      rowLeft3: 100,
      rowTotal: 1000000,
      // 日期
      currentDate: '',
      // 时间
      currentTime: '',

      pcount: 396,// 累计项目数
      housesCount: 227900,// 累计房源数
      currentHouses: 227900,// 当前持有房源数
      operableHouses: 227900,// 可运营房源数
      rentalRate: 50,// 出租率
      rentalPrice: 2750,// 出租均价
      serviceEnt: 5910,// 累计服务企业数
      serviceTenant: 366400,// 累计服务租客数
      smartDevices: 200000,// 接入智能设备数
      cooperativeOp: 5000,// 合作运营商数
      allianceEnt: 2100,// 入库联盟企业数

      aroundHide: true
    }
  },
  created() {
  },
  mounted() {
    this.list();
    // this.mapChart();
    this.totalProgress();
    this.timeStamp();
    // this.row()
  },
  methods: {
    // map初始化
    mapChart() {
      var mapName = 'china'
      var data = [
        { name: "北京", value: 199 },
        { name: "天津", value: 42 },
        { name: "河北", value: 102 },
        { name: "山西", value: 81 },
        { name: "内蒙古", value: 47 },
        { name: "辽宁", value: 67 },
        { name: "吉林", value: 82 },
        { name: "黑龙江", value: 123 },
        { name: "上海", value: 24 },
        { name: "江苏", value: 92 },
        { name: "浙江", value: 114 },
        { name: "安徽", value: 109 },
        { name: "福建", value: 116 },
        { name: "江西", value: 91 },
        { name: "山东", value: 119 },
        { name: "河南", value: 137 },
        { name: "湖北", value: 116 },
        { name: "湖南", value: 114 },
        { name: "重庆", value: 91 },
        { name: "四川", value: 125 },
        { name: "贵州", value: 62 },
        { name: "云南", value: 83 },
        { name: "西藏", value: 9 },
        { name: "陕西", value: 80 },
        { name: "甘肃", value: 56 },
        { name: "青海", value: 10 },
        { name: "宁夏", value: 18 },
        { name: "新疆", value: 180 },
        { name: "广东", value: 123 },
        { name: "广西", value: 59 },
        { name: "海南", value: 14 },
      ];

      var geoCoordMap = {};
      var toolTipData = [
        { name: "北京", value: [{ name: "科技人才总数", value: 95 }, { name: "理科", value: 82 }] },
        { name: "天津", value: [{ name: "文科", value: 22 }, { name: "理科", value: 20 }] },
        { name: "河北", value: [{ name: "文科", value: 60 }, { name: "理科", value: 42 }] },
        { name: "山西", value: [{ name: "文科", value: 40 }, { name: "理科", value: 41 }] },
        { name: "内蒙古", value: [{ name: "文科", value: 23 }, { name: "理科", value: 24 }] },
        { name: "辽宁", value: [{ name: "文科", value: 39 }, { name: "理科", value: 28 }] },
        { name: "吉林", value: [{ name: "文科", value: 41 }, { name: "理科", value: 41 }] },
        { name: "黑龙江", value: [{ name: "文科", value: 35 }, { name: "理科", value: 31 }] },
        { name: "上海", value: [{ name: "文科", value: 12 }, { name: "理科", value: 12 }] },
        { name: "江苏", value: [{ name: "文科", value: 47 }, { name: "理科", value: 45 }] },
        { name: "浙江", value: [{ name: "文科", value: 57 }, { name: "理科", value: 57 }] },
        { name: "安徽", value: [{ name: "文科", value: 57 }, { name: "理科", value: 52 }] },
        { name: "福建", value: [{ name: "文科", value: 59 }, { name: "理科", value: 57 }] },
        { name: "江西", value: [{ name: "文科", value: 49 }, { name: "理科", value: 42 }] },
        { name: "山东", value: [{ name: "文科", value: 67 }, { name: "理科", value: 52 }] },
        { name: "河南", value: [{ name: "文科", value: 69 }, { name: "理科", value: 68 }] },
        { name: "湖北", value: [{ name: "文科", value: 60 }, { name: "理科", value: 56 }] },
        { name: "湖南", value: [{ name: "文科", value: 62 }, { name: "理科", value: 52 }] },
        { name: "重庆", value: [{ name: "文科", value: 47 }, { name: "理科", value: 44 }] },
        { name: "四川", value: [{ name: "文科", value: 65 }, { name: "理科", value: 60 }] },
        { name: "贵州", value: [{ name: "文科", value: 32 }, { name: "理科", value: 30 }] },
        { name: "云南", value: [{ name: "文科", value: 42 }, { name: "理科", value: 41 }] },
        { name: "西藏", value: [{ name: "文科", value: 5 }, { name: "理科", value: 4 }] },
        { name: "陕西", value: [{ name: "文科", value: 38 }, { name: "理科", value: 42 }] },
        { name: "甘肃", value: [{ name: "文科", value: 28 }, { name: "理科", value: 28 }] },
        { name: "青海", value: [{ name: "文科", value: 5 }, { name: "理科", value: 5 }] },
        { name: "宁夏", value: [{ name: "文科", value: 10 }, { name: "理科", value: 8 }] },
        { name: "新疆", value: [{ name: "文科", value: 36 }, { name: "理科", value: 31 }] },
        { name: "广东", value: [{ name: "文科", value: 63 }, { name: "理科", value: 60 }] },
        { name: "广西", value: [{ name: "文科", value: 29 }, { name: "理科", value: 30 }] },
        { name: "海南", value: [{ name: "文科", value: 8 }, { name: "理科", value: 6 }] },
      ];

      /*获取地图数据*/
      var mapFeatures = this.$echarts.getMap(mapName).geoJson.features;
      mapFeatures.forEach(function(v) {
        // 地区名称
        var name = v.properties.name;
        // 地区经纬度
        geoCoordMap[name] = v.properties.cp;

      });

      // console.log(data)
      // console.log(toolTipData)
      var max = 480,
        min = 9; // todo 
      var maxSize4Pin = 100,
        minSize4Pin = 20;

      var convertData = function(data) {
        var res = [];
        for (var i = 0; i < data.length; i++) {
          var geoCoord = geoCoordMap[data[i].name];
          if (geoCoord) {
            res.push({
              name: data[i].name,
              value: geoCoord.concat(data[i].value),
            });
          }
        }
        return res;
      };
      const option = {
        tooltip: {
          padding: 0,
          enterable: true,
          transitionDuration: 1,
          textStyle: {
            color: '#000',
            decoration: 'none',
          },
          // position: function (point, params, dom, rect, size) {
          //   return [point[0], point[1]];
          // },
          formatter: function(params) {
            // console.log(params)
            var tipHtml = '';
            tipHtml = '<div style="width:280px;height:180px;background:rgba(22,80,158,0.8);border:1px solid rgba(7,166,255,0.7)">'
              + '<div style="width:100%;height:40px;line-height:40px;border-bottom:2px solid rgba(7,166,255,0.7);padding:0 20px">' + '<i style="display:inline-block;width:8px;height:8px;background:#16d6ff;border-radius:40px;">' + '</i>'
              + '<span style="margin-left:10px;color:#fff;font-size:16px;">' + params.name + '</span>' + '</div>'
              + '<div style="padding:20px">'
              + '<p style="color:#fff;font-size:12px;">' + '<i style="display:inline-block;width:10px;height:10px;background:#16d6ff;border-radius:40px;margin:0 8px">' + '</i>'
              + '单位总数：' + '<span style="color:#11ee7d;margin:0 6px;">' + toolTipData.length + '</span>' + '个' + '</p>'
              + '<p style="color:#fff;font-size:12px;">' + '<i style="display:inline-block;width:10px;height:10px;background:#16d6ff;border-radius:40px;margin:0 8px">' + '</i>'
              + '总人数：' + '<span style="color:#f48225;margin:0 6px;">' + toolTipData.length + '</span>' + '个' + '</p>'
              + '<p style="color:#fff;font-size:12px;">' + '<i style="display:inline-block;width:10px;height:10px;background:#16d6ff;border-radius:40px;margin:0 8px">' + '</i>'
              + '总人数：' + '<span style="color:#f4e925;margin:0 6px;">' + toolTipData.length + '</span>' + '个' + '</p>'
              + '<p style="color:#fff;font-size:12px;">' + '<i style="display:inline-block;width:10px;height:10px;background:#16d6ff;border-radius:40px;margin:0 8px">' + '</i>'
              + '总人数：' + '<span style="color:#25f4f2;margin:0 6px;">' + toolTipData.length + '</span>' + '个' + '</p>'
              + '</div>' + '</div>';
            return tipHtml;
          }

        },

        visualMap: {
          show: true,
          min: 0,
          max: 200,
          left: '10%',
          top: 'bottom',
          calculable: true,
          seriesIndex: [1],
          inRange: {
            color: ['#04387b', '#467bc0'] // 蓝绿
          }
        },
        geo: {
          show: true,
          map: mapName,
          label: {
            normal: {
              show: false
            },
            emphasis: {
              show: false,
            }
          },
          roam: false,
          itemStyle: {
            normal: {
              areaColor: '#023677',
              borderColor: '#1180c7',
            },
            emphasis: {
              areaColor: '#4499d0',
            }
          }
        },
        series: [{
          name: '散点',
          type: 'scatter',
          coordinateSystem: 'geo',
          data: convertData(data),
          symbolSize: function(val) {
            return val[2] / 10;
          },
          label: {
            normal: {
              formatter: '{b}',
              position: 'right',
              show: true
            },
            emphasis: {
              show: true
            }
          },
          itemStyle: {
            normal: {
              color: '#fff'
            }
          }
        },
        {
          type: 'map',
          map: mapName,
          geoIndex: 0,
          aspectScale: 0.75, //长宽比
          showLegendSymbol: false, // 存在legend时显示
          label: {
            normal: {
              show: true
            },
            emphasis: {
              show: false,
              textStyle: {
                color: '#fff'
              }
            }
          },
          roam: true,
          itemStyle: {
            normal: {
              areaColor: '#031525',
              borderColor: '#3B5077',
            },
            emphasis: {
              areaColor: '#2B91B7'
            }
          },
          animation: false,
          data: data
        },
        {
          name: '点',
          type: 'scatter',
          coordinateSystem: 'geo',
          zlevel: 6,
        },
        {
          name: 'Top 5',
          type: 'effectScatter',
          coordinateSystem: 'geo',
          data: convertData(data.sort(function(a, b) {
            return b.value - a.value;
          }).slice(0, 10)),
          symbolSize: function(val) {
            return val[2] / 10;
          },
          showEffectOn: 'render',
          rippleEffect: {
            brushType: 'stroke'
          },
          hoverAnimation: true,
          label: {
            normal: {
              formatter: '{b}',
              position: 'left',
              show: false
            }
          },
          itemStyle: {
            normal: {
              color: 'yellow',
              shadowBlur: 10,
              shadowColor: 'yellow'
            }
          },
          zlevel: 1
        },

        ]
      };
      init('.map_chart', option);
    },
    // 计算进度
    barProgress() {
      return `bottom: calc(-125% + ${this.rentalRate}%);`
    },
    // total进度
    totalProgress() {
      const option = {
        color: [
          {
            type: "linear",
            x: 0,
            y: 0,
            x2: 1,
            y2: 0,
            colorStops: [
              {
                offset: 0,
                color: "rgba(238, 245, 125, 1)"
              },
              {
                offset: 0.5,
                color: "rgba(238, 245, 125, 1)"
              },
              {
                offset: 1,
                color: "rgba(238, 245, 125, 1)"
              }
            ],
            global: false
          }
        ],
        angleAxis: {
          max: 100,
          show: false,
          startAngle: 180,
        },
        radiusAxis: {
          type: "category",
          show: false
        },
        polar: {
          radius: '190%',
          center: ["50%", "50%"]
        },
        series: [
          {
            type: "bar",
            roundCap: true,
            barWidth: 5,
            showBackground: true,
            backgroundStyle: {
              color: "rgba(0,0,0,0)"
            },
            coordinateSystem: "polar",
            name: "库存周转率",
            data: [90]
          }
        ],
        animationDuration: 4000,
        animaationEasing: 'cubicInOut',
      }
      init('.total_progress', option)
    },
    // row进度
    rowProgress(num, count) {
      // setTimeout(() => {
      //   this.rowLeft1 = 100 - (count / this.rowTotal) * 100
      //   console.log(this.rowLeft1);
      // }, 200);
      if (num === 1) {
        setTimeout(() => {
          this.rowLeft1 = 100 - (count / this.rowTotal) * 100
        }, 200);
      } else if (num === 2) {
        setTimeout(() => {
          this.rowLeft2 = 100 - (count / this.rowTotal) * 100
        }, 200);
      } else if (num === 3) {
        setTimeout(() => {
          this.rowLeft3 = 100 - (count / this.rowTotal) * 100
        }, 200);
      }
      // return `left: calc(-${this.rowLeft}% + 0.06rem);`
      // timer = setInterval(loading, 1000)
    },
    row() {
      let timer = setInterval(() => {
        this.rowLeft = this.rowLeft - 1;
        if (this.rowLeft <= 0) {
          clearInterval(timer)
        }
      }, 100);
    },
    timeStamp() {
      const time = () => {
        let dt = new Date();
        let y = dt.getFullYear();
        let mt = dt.getMonth() + 1;
        let day = dt.getDate();
        let h = dt.getHours();//获取时
        let m = dt.getMinutes();//获取分
        let s = dt.getSeconds();//获取秒
        if (mt < 10) {
          mt = `0${mt}`;
        };
        if (day < 10) {
          day = `0${day}`;
        };
        if (h < 10) {
          h = `0${h}`;
        };
        if (m < 10) {
          m = `0${m}`
        };
        if (s < 10) {
          s = `0${s}`
        };
        this.currentDate = `${y}.${mt}.${day}`;
        this.currentTime = `${h}:${m}:${s}`;
      }
      setInterval(time, 1000)

    },
    // 获取所有数据
    list() {
      panoramaList().then((res) => {

        console.log(res.data.data.data[0]);
        // { allianceEnt, cooperativeOp, currentHouses, housesCount, operableHouses, pcount, rentalPrice, serviceEnt, serviceTenant, smartDevices } = res.data.data.data[0];
        this.allianceEnt = res.data.data.data[0].allianceEnt;
        this.cooperativeOp = res.data.data.data[0].cooperativeOp;
        this.currentHouses = res.data.data.data[0].currentHouses;
        this.housesCount = res.data.data.data[0].housesCount;
        this.operableHouses = res.data.data.data[0].operableHouses;
        this.pcount = res.data.data.data[0].pcount;
        this.rentalPrice = res.data.data.data[0].rentalPrice;
        this.rentalRate = res.data.data.data[0].rentalRate;
        this.serviceEnt = res.data.data.data[0].serviceEnt;
        this.serviceTenant = res.data.data.data[0].serviceTenant;
        this.smartDevices = res.data.data.data[0].smartDevices;
        this.rowProgress(1, this.housesCount);
        this.rowProgress(2, this.currentHouses);
        this.rowProgress(3, this.operableHouses);
      });
    },
    largeValue(val) {
      // 10W
      if (val > 100000 && val <= 1000000) {
        let num1 = val.toString().slice(0, 2);
        let num2 = val.toString().slice(2, 4);
        return `${num1}.${num2}`
      } else if (val > 1000000) {
        let num1 = val.toString().slice(0, 3);
        let num2 = val.toString().slice(3, 5);
        return `${num1}.${num2}`
      }
    }
  },
  computed: {
    housesCountCal() {
      return this.largeValue(this.housesCount)
    },
    currentHousesCal() {
      return this.largeValue(this.currentHouses)
    },
    operableHousesCal() {
      return this.largeValue(this.operableHouses)
    },
    serviceTenantCal() {
      return this.largeValue(this.serviceTenant)
    },
    smartDevicesCal() {
      return this.largeValue(this.smartDevices)
    }

  },
  watch: {
    '$store.state.region.aroundHide'(next) {
      this.aroundHide = next
    }
  }
};
</script>
<style scoped lang="less">
.panorama {
  height: 13.5rem;
  width: 24rem;
  background: url(../assets/images/背景.png) no-repeat center/100%;
  color: aliceblue;
  position: relative;
  font-family: 'lt';
  overflow: hidden;
}

.header {
  height: 2.58rem;

  .top {
    height: 0.85rem;
    background: url(../assets/images/顶部标题.png) no-repeat;
    background-size: 100% 100%;
    position: relative;
    font-size: 0.2rem;

    .top_left {
      position: absolute;
      width: 7.3063rem;
      height: 0.53rem;
      top: 0.18rem;
      left: 0;

      .back {
        position: absolute;
        left: 0.4rem;
        top: 0.25rem;
        width: 0.6875rem;
        height: 0.25rem;
        background: url(../assets/images/返回.png) no-repeat;
        background-size: 100% 100%;
      }

      .content {
        position: absolute;
        left: 1.5625rem;
        top: 0.1rem;
        height: 0.16rem;
        font-size: 0.16rem;
        line-height: 0.55rem;
        font-weight: 400;
      }
    }

    h1 {
      font-size: 0.5rem;
      text-align: center;
      line-height: 0.75rem;
      font-family: 'fzlt';
      font-weight: 400;
      // 设置行间距
      letter-spacing: .0625rem;
      line-height: 1.8;
    }

    .top_right {
      position: absolute;
      top: 0.2rem;
      right: 0;
      width: 7.28rem;
      height: 0.52rem;

      .date {
        position: absolute;
        top: 0.12rem;
        right: 1.56rem;
        height: 0.16rem;
        line-height: 0.43rem;
        font-size: 0.16rem;
        opacity: 0.7;
      }

      .time {
        position: absolute;
        width: 1rem;
        height: 0.46rem;
        top: 0.1rem;
        right: 0.4rem;
        font-size: 0.23rem;
        font-weight: 400;
        line-height: 0.43rem;
      }
    }
  }

  // .select {
  //   background: url(../assets/images/区域选择/矩形.png) no-repeat;
  //   background-size: 100% 100%;
  //   width: 2.56rem;
  //   height: 0.45rem;
  //   position: absolute;
  //   top: 1.13rem;
  //   left: 0.41rem;
  //   line-height: 0.45rem;
  //   padding-left: 0.21rem;

  //   .select_icon {
  //     background: url(../assets/images/区域选择/下拉.png) no-repeat;
  //     background-size: 100% 100%;
  //     width: 0.18rem;
  //     height: 0.09rem;
  //     position: absolute;
  //     top: 0.18rem;
  //     right: 0.17rem;
  //   }
  // }
}

.main {
  display: flex;
  margin: 0 auto;
  justify-content: center;
  height: 10.93rem;
  padding: 0 0.66rem;
  position: relative;

  .child {
    overflow: hidden;

    &:nth-child(1) {
      width: 4.2rem;
    }

    // &:nth-child(2) {
    //   flex: 1;
    // }

    &:nth-child(3) {
      width: 4.2rem;
      display: flex;
      flex-flow: column;
      align-items: flex-end;
    }
  }

  .main_left {
    position: absolute;
    left: 0.66rem;
    top: 0;
    animation-duration: 1s;
  }

  .main_right {
    position: absolute;
    right: 0.66rem;
    top: 0;
    animation-duration: 1s;
  }

  .main_middle {
    width: 15rem;
    height: 10rem;
  }

  // 累计数
  .total_data {
    position: relative;
    margin-top: 0.48rem;
    background: url(../assets/images/显示数据/项目数.png) no-repeat;
    background-size: 100% 100%;
    height: 1.8625rem;
    width: 4.1125rem;

    .total_left {
      position: absolute;
      left: 0.25rem;
      top: 0.24rem;
      width: 1.35rem;
      height: 1.35rem;
      border-radius: 50%;

      .project_icon {
        background: url(../assets/images/显示数据/项目数/项目.png) no-repeat;
        background-size: 100% 100%;
        width: 0.55rem;
        height: 0.63rem;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
      }

      .total_progress {
        position: absolute;
        width: 100%;
        height: 1.35rem;
        top: 50%;
        left: 50%;
        transform: translate(-49%, -49%);
      }
    }

    .total_right {
      position: absolute;
      top: 0;
      right: 0.25rem;
      width: 1.81rem;
      height: 1.1rem;
      display: flex;
      flex-flow: column;
      justify-content: space-between;

      .r_top {
        height: 0.54rem;
        color: #eef57d;
        font-size: 0.63rem;
        font-weight: 400;
        line-height: 2;
      }

      .r_bottom {
        height: 0.3rem;
        font-size: 0.32rem;
        font-weight: 400;
        color: #80c2ff;
        line-height: 1.11rem;
      }
    }
  }

  // 项目规模
  .project_scale {
    width: 4.11rem;
    height: 4.16rem;
    margin-top: 2.96rem;

    .p_header {
      background: url(../assets/images/标题.png) no-repeat;
      background-size: 100% 100%;
      height: 0.6rem;
      position: relative;

      .p_icon {
        background: url(../assets/images/1房源项目规模/项目规模.png) no-repeat;
        background-size: 100% 100%;
        position: absolute;
        width: 0.3125rem;
        height: 0.3rem;
        top: 0.24rem;
        left: 0.09rem;
      }

      .p_title {
        position: absolute;
        left: 0.67rem;
        top: 0;
        height: 0.36rem;
        font-size: 0.38rem;
        font-weight: 400;
        color: #00deff;
        line-height: 0.4rem;
      }
    }

    .p_main {
      display: flex;
      flex-flow: column;
      justify-content: space-between;
      height: 2.99rem;
      margin-top: 0.57rem;

      .row {
        width: 4.11rem;

        .row_info {
          display: flex;
          justify-content: space-between;
          align-items: flex-end;

          .info_left {
            display: flex;
            align-items: flex-end;

            .info_icon {

              &.icon1 {
                background: url('../assets/images/1房源项目规模/累计房源.png') no-repeat center bottom/100%;
                width: .2125rem;
                height: .2625rem;
              }

              &.icon2 {
                background: url('../assets/images/1房源项目规模/项目规模.png') no-repeat center bottom/100%;
                width: .225rem;
                height: .2625rem;
              }

              &.icon3 {
                background: url('../assets/images/1房源项目规模/运营房源.png') no-repeat center bottom/100%;
                width: .2625rem;
                height: .25rem;
              }
            }

            .info_title {
              font-size: 0.22rem;
              font-weight: 400;
              line-height: 1;
              color: #9FDCFF;
              padding-left: 0.15rem;
            }
          }

          .info_right {
            font-size: 0.3rem;
            font-weight: 400;
            color: #fff;
            line-height: 1;
          }
        }

        .row_progress {
          margin-top: 0.16rem;
          height: 0.25rem;
          padding: 0 0.06rem;
          position: relative;
          overflow: hidden;

          &.progress1 {
            background: url('../assets/images/1房源项目规模/1.png') no-repeat center/100%;
          }

          &.progress2 {
            background: url('../assets/images/1房源项目规模/2.png') no-repeat center/100%;
          }

          &.progress3 {
            background: url('../assets/images/1房源项目规模/3.png') no-repeat center/100%;
          }

          img {
            position: absolute;
            width: 4rem;
            height: 0.15rem;
            top: 50%;
            transform: translateY(-50%);
            transition: left 3s;
          }

        }
      }
    }
  }

  // 运营效率
  .run_efficiency {
    width: 4.0625rem;
    height: 5.19rem;
    margin-top: 0.06rem;

    .r_header {
      background: url('../assets/images/标题.png') no-repeat center/100%;
      height: 0.6rem;
      position: relative;

      img {
        position: absolute;
        top: 0.25rem;
        left: 0.08rem;
        width: 0.34rem;
        height: 0.29rem;
      }

      .r_title {
        position: absolute;
        left: 0.67rem;
        top: 0;
        height: 0.36rem;
        font-size: 0.38rem;
        font-weight: 400;
        color: #00deff;
        line-height: 0.4rem;
      }
    }

    .r_main {
      margin-top: 0.55rem;
      width: 4.07rem;

      .rm_ripple {
        display: flex;
        justify-content: space-between;

        .ripple {

          .ripple_bar {
            background: url('../assets/images/2运营效率/涟漪底色.png') no-repeat center/100%;
            width: 1.75rem;
            height: 1.75rem;
            position: relative;
            overflow: hidden;

            .bar_box {
              position: absolute;
              width: 1.46rem;
              height: 1.46rem;
              top: 0.15rem;
              left: 0.14rem;
              border-radius: 50%;
              overflow: hidden;

              .bar_inner {
                width: 100%;
                height: 100%;
                position: absolute;
                left: 0;
                // bottom: calc(-125% + 100%);
                background: #00FFFF;

                .box_waves {
                  position: absolute;
                  left: 0;
                  bottom: 100%;
                  width: 200%;
                  stroke: none;

                  &:nth-child(1) {
                    fill: #00FFFF;
                    transform: translate(-50%, 0);
                    z-index: 3;
                    animation: wave-move1 1.5s linear infinite;
                    /* svg重合有一条线 */
                    margin-bottom: -2px;
                  }

                  &:nth-child(2) {
                    fill: rgba(40, 187, 255, 0.5);
                    transform: translate(0, 0);
                    z-index: 2;
                    animation: wave-move2 3s linear infinite;
                  }

                  &:nth-child(3) {
                    fill: #2084cc;
                    transform: translate(-50%, 0);
                    z-index: 1;
                    animation: wave-move1 3s linear infinite;
                  }
                }

                @keyframes wave-move1 {
                  100% {
                    transform: translate(0, 0);
                  }
                }

                @keyframes wave-move2 {
                  100% {
                    transform: translate(-50%, 0);
                  }
                }
              }
            }


            .ripple_num {
              position: absolute;
              left: 50%;
              top: 50%;
              transform: translate(-50%, -50%);
              font-size: 0.3rem;
              height: 0.54rem;
              font-weight: 400;
              color: #fff;
              line-height: 1.2;
              text-align: center;
              z-index: 5;
            }
          }

          &:nth-child(2) {
            .bar_box {
              .bar_inner {
                background: #31F5A9;

                .box_waves {
                  &:nth-child(1) {
                    fill: #31F5A9;
                  }
                }
              }
            }

            .ripple_num {
              div {
                font-size: 0.2rem;
              }
            }

            .ripple_tag {
              padding-left: 0.48rem;
            }
          }

          .ripple_tag {
            padding: 0.23rem 0 0 0.57rem;
            font-size: 0.2rem;
            font-weight: 400;
            color: #9fdcff;
            line-height: 1;
          }
        }

      }

      .rm_group {
        background: url('../assets/images/2运营效率/组 127.png') no-repeat center/100%;
        height: 1.3875rem;
        margin-top: 0.48rem;
        position: relative;

        .group {
          position: absolute;
          display: flex;
          align-items: center;
          justify-content: space-between;

          .group_info {
            .gi_top {
              font-size: 0.3rem;
              line-height: 1;
              font-weight: 400;

              span {
                font-size: 0.2rem;
              }
            }

            .gi_bottom {
              font-size: 0.2rem;
              line-height: 2;
              font-weight: 400;
              color: #9FDCFF;
            }
          }

          &:nth-child(1) {
            .group_info {
              padding-right: 0.16rem;
            }

            img {
              width: .2875rem;
              height: .2375rem;
              // padding-right: 0.16rem;

            }
          }

          &:nth-child(2) {
            .group_info {
              padding-left: 0.16rem;
            }

            img {
              width: .3rem;
              height: .25rem;
            }
          }
        }

        .group_left {
          left: 0;
          top: 0.1rem;
        }

        .group_right {
          top: 0.58rem;
          right: 0;
        }

      }
    }

  }

  // 产业生态
  .product_ecology {
    width: 4.0625rem;
    height: 3.86rem;
    margin-top: 0.89rem;

    .pe_header {
      background: url('../assets/images/标题.png') no-repeat center/100%;
      height: 0.6rem;
      position: relative;

      img {
        position: absolute;
        top: 0.25rem;
        left: 0.09rem;
        width: .325rem;
        height: .2625rem;
      }

      .pe_title {
        position: absolute;
        left: 0.67rem;
        top: 0;
        height: 0.36rem;
        font-size: 0.38rem;
        font-weight: 400;
        color: #00deff;
        line-height: 1;
      }
    }

    .pe_main {
      margin-top: 0.57rem;
      height: 2.69rem;
      position: relative;

      .pe_info {
        position: absolute;
        border-left: 0.04rem solid;
        // display: flex;
        // flex-flow: column;
        // align-items: flex-start;
        z-index: 7;
        padding-left: 0.11rem;

        .pi_top {
          font-size: 0.3rem;
          line-height: 1;
          font-weight: 400;

          span {
            font-size: 0.2rem;
          }
        }

        .pi_bottom {
          font-size: 0.2rem;
          line-height: 2;
          font-weight: 400;
          color: #9FDCFF;

        }
      }

      .pe_chart {
        width: 2.30rem;
        height: 2.30rem;
        position: absolute;
        left: 0.8rem;
        top: -0.2rem;

        .pe_mask {
          width: 1.17rem;
          height: 0.95rem;
          position: absolute;
          left: -0.1rem;
          top: 0.2rem;
          background-color: #002561;
          z-index: 5;
        }

        // background-color: lightpink;

        .pe_chart_1 {
          position: absolute;
          left: 50%;
          top: 50%;
          transform: translate(-50%, -50%);
          width: 1.50rem;
          height: 1.50rem;
          border-radius: 50%;
          box-sizing: border-box;
          z-index: 4;

          .inner_1 {
            position: absolute;
            top: -0.075rem;
            left: -0.075rem;
            bottom: -0.075rem;
            right: -0.075rem;
            border-radius: 50%;
            background: conic-gradient(#EEF57D 0%, #EEF57D 75%, transparent 75%);
            // animation: rotate 3s ease-in-out infinite;
            transform: rotate(0deg);
            // transition-origin: 50% 50%;

            &::before {
              content: "";
              position: absolute;
              top: .0625rem;
              left: .0625rem;
              bottom: .0625rem;
              right: .0625rem;
              background: #002561;
              border-radius: 50%;
              // z-index: 1;
            }

            &::after {
              content: "";
              position: absolute;
              width: 0.15rem;
              height: 0.15rem;
              top: 0;
              left: 50%;
              transform: translate(-50%, -3px);
              border-radius: 50%;
              background: #EEF57D;
              z-index: 9;
            }
          }
        }

        .pe_chart_2 {
          position: absolute;
          left: 50%;
          top: 50%;
          transform: translate(-50%, -50%);
          width: 1.75rem;
          height: 1.75rem;
          border-radius: 50%;
          box-sizing: border-box;
          z-index: 3;

          .inner_2 {
            position: absolute;
            top: -0.075rem;
            left: -0.075rem;
            bottom: -0.075rem;
            right: -0.075rem;
            border-radius: 50%;
            background: conic-gradient(#31F5A9 0%, #31F5A9 50%, transparent 50%);
            // animation: rotate 3s ease-in-out infinite;
            transform: rotate(90deg);
            // transition-origin: 50% 50%;

            &::before {
              content: "";
              position: absolute;
              top: .0625rem;
              left: .0625rem;
              bottom: .0625rem;
              right: .0625rem;
              background: #002561;
              border-radius: 50%;
            }

            &::after {
              content: "";
              position: absolute;
              width: 0.15rem;
              height: 0.15rem;
              top: 0;
              left: 50%;
              transform: translate(-50%, -3px);
              border-radius: 50%;
              background: #31F5A9;
            }
          }
        }

        .pe_chart_3 {
          position: absolute;
          left: 50%;
          top: 50%;
          transform: translate(-50%, -50%);
          width: 2rem;
          height: 2rem;
          border-radius: 50%;
          box-sizing: border-box;

          .inner_3 {
            position: absolute;
            top: -0.075rem;
            left: -0.075rem;
            bottom: -0.075rem;
            right: -0.075rem;
            border-radius: 50%;
            background: conic-gradient(#00FFFF 0%, #00FFFF 25%, transparent 25%);
            // animation: rotate 3s ease-in-out infinite;
            transform: rotate(180deg);
            // transition-origin: 50% 50%;

            &::before {
              content: "";
              position: absolute;
              top: .0625rem;
              left: .0625rem;
              bottom: .0625rem;
              right: .0625rem;
              background: #002561;
              border-radius: 50%;
            }

            &::after {
              content: "";
              position: absolute;
              width: 0.15rem;
              height: 0.15rem;
              top: 0;
              left: 50%;
              transform: translate(-50%, -3px);
              border-radius: 50%;
              background: #00FFFF;
            }
          }
        }



      }

      .pe_left {
        top: 0;
        left: 0;
        border-left-color: #EEF57D
      }

      .pe_right {
        top: 0;
        right: 0;
        border-left-color: #31F5A9;
      }

      .pe_bottom {
        bottom: 0;
        right: 0.24rem;
        border-left-color: #00FFFF;
      }
    }
  }

  // .main_left {}

  // .main_middle {}

  // .main_right {}
  .map_chart {
    width: 100%;
    height: 100%;
  }
}

.footer {
  background: url('../assets/images/底部按钮/组 124.png') no-repeat center/100%;
  // width: 5.0625rem;
  height: .4875rem;
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  bottom: 0;
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  background-size: 105%;

  .footer_focus {
    width: 2.425rem;
    height: .4rem;
    text-align: center;
    font-weight: 400;
    font-size: 0.2rem;
    line-height: 1.8;
  }

  .footer_left {
    background: url('../assets/images/底部按钮/组 125.png') no-repeat center/100%;
    margin-right: .025rem;
  }

  .footer_right {
    background: url('../assets/images/底部按钮/组 126.png') no-repeat center/100%;
  }
}
</style>
