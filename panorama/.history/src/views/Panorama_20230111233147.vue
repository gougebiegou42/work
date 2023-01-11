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
          <div class="date">2022.11.7 星期一</div>
          <div class="time">11:27:28</div>
        </div>
      </div>
      <div class="select">
        全国
        <div class="select_icon"></div>
      </div>
    </div>
    <!-- 主体内容 -->
    <div class="main">
      <div class="child main_left">
        <!-- 累计项目数 -->
        <div class="total_data">
          <div class="total_left">
            <div class="project_icon"></div>
          </div>
          <div class="total_right">
            <div class="r_top">395个</div>
            <div class="r_bottom">累计项目数</div>
          </div>
          <div class="chart"></div>
        </div>
        <!-- 房源项目规模 -->
        <div class="project_scale">
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
                <div class="info_right">22.79 万间</div>
              </div>
              <div class="row_progress progress1">
                <img src="../assets/images/1房源项目规模/黄.png">
              </div>
            </div>
            <div class="row pm_2">
              <div class="row_info">
                <div class="info_left">
                  <div class="info_icon icon2"></div>
                  <div class="info_title">累计房源数</div>
                </div>
                <div class="info_right">22.79 万间</div>
              </div>
              <div class="row_progress progress2">
                <img src="../assets/images/1房源项目规模/青.png">
              </div>
            </div>
            <div class="row pm_3">
              <div class="row_info">
                <div class="info_left">
                  <div class="info_icon icon3"></div>
                  <div class="info_title">累计房源数</div>
                </div>
                <div class="info_right">22.79 万间</div>
              </div>
              <div class="row_progress progress3">
                <img src="../assets/images/1房源项目规模/绿.png">
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="child main_middle">
        <!-- 地图 -->
        <div class="chart"></div>
      </div>
      <div class="child main_right">
        <!-- 运营效率 -->
        <div class="run_efficiency"></div>
        <!-- 产业生态 -->
        <div class="product_ecology"></div>
      </div>
    </div>
  </div>
</template>

<script>
// import * as echarts from 'echarts';

export default {
  mounted() {
    const main = document.querySelector('.main');
    console.log(main);
  },
  // const myChart = echarts.init(main);
};
</script>
<style scoped lang="less">
.panorama {
  height: 13.5rem;
  width: 24rem;
  background: url(../assets/images/背景.png) no-repeat top center;
  color: aliceblue;
  position: relative;
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
      // 设置行间距
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

  .select {
    background: url(../assets/images/区域选择/矩形.png) no-repeat;
    background-size: 100% 100%;
    width: 2.56rem;
    height: 0.45rem;
    position: absolute;
    top: 1.13rem;
    left: 0.41rem;
    line-height: 0.45rem;
    padding-left: 0.21rem;

    .select_icon {
      background: url(../assets/images/区域选择/下拉.png) no-repeat;
      background-size: 100% 100%;
      width: 0.18rem;
      height: 0.09rem;
      position: absolute;
      top: 0.18rem;
      right: 0.17rem;
    }
  }
}

.main {
  display: flex;
  min-width: 1024px;
  max-width: 1920px;
  margin: 0 auto;
  padding: 0 0.66rem;

  .child {
    flex: 3;
    overflow: hidden;

    &:nth-child(2) {
      flex: 5;
    }
  }

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
      height: 1.34rem;
      display: flex;
      align-items: center;
      justify-content: center;

      .project_icon {
        background: url(../assets/images/显示数据/项目数/项目.png) no-repeat;
        background-size: 100% 100%;
        width: 0.55rem;
        height: 0.63rem;
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
        line-height: 1.11rem;
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

  .project_scale {
    width: 4.11rem;
    height: 4.16rem;
    margin-top: 2.96rem;

    .p_header {
      background: url(../assets/images/1房源项目规模/标题.png) no-repeat;
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

            .info_icon {

              &.icon1 {
                background: url('../assets/images/1房源项目规模/累计房源.png') no-repeat center/100%;
                width: .225rem;
                height: .2625rem;
              }

              &.icon2 {
                background: url('../assets/images/1房源项目规模/项目规模.png') no-repeat center/100%;
                width: .225rem;
                height: .2625rem;
              }

              &.icon3 {
                background: url('../assets/images/1房源项目规模/运营房源.png') no-repeat center/100%;
                width: .2625rem;
                height: .25rem;
              }
            }

            .info_title {
              font-size: 0.22rem;
              font-weight: 400;
              line-height: 0.3rem;
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
          position: relative;

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
            left: calc(0% + 0.06rem);
          }
        }
      }
    }
  }

  // .main_left {}

  // .main_middle {}

  // .main_right {}
}
</style>
