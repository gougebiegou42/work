<template>
  <div class="box_prefecture" v-show="hide">
    <div class="main_left animate__fadeInLeft">
      <div class="left_item" v-for="(item, index) in leftData" :key="index" @click="activeClick(index)">
        <img class="pic" :src="shiftImg('left', isActive == index ? item.img + '2' : item.img)" alt="">
        <div class="title" :class="isActive == index ? 'active' : ''">{{
          item.title
        }}</div>
      </div>
    </div>
    <div class="main_right animate__fadeInRight" v-show="isActive === 0">
      <!-- 累计项目数 -->
      <div class="total_data">
        <div class="td_left">
          <img class="l_top" src="../assets/images/regionRight/累计项目.png" alt="">
          <div class="l_bottom">累计项目数</div>
        </div>
        <div class="td_right">66<span>个</span></div>
      </div>
      <!-- 房源项目规模 -->
      <div class="project_scale">
        <div class="ps_item" v-for="(item, index) in projectScaleData" :key="index">
          <img class="project_left" :src="shiftImg('right', item.leftImg)" alt="">
          <div class="project_right">
            <div class="r_title">{{ item.projectTitle }}</div>
            <div class="r_num">{{ item.projectNum }}<span>万间</span></div>
          </div>
        </div>
      </div>
    </div>
    <div class="main_right left3 animate__fadeInRight" v-show="isActive === 2">
      <!-- 产业生态 -->
      <div class="project_scale">
        <div class="ps_item" v-for="(item, index) in productEcology" :key="index">
          <img class="project_left" :src="shiftImg('ecology', item.leftImg)" alt="">
          <div class="project_right">
            <div class="r_title">{{ item.projectTitle }}</div>
            <div class="r_num">{{ item.projectNum }}<span>{{ index=== 0 ? '家' : '万台' }}</span></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      leftData: [
        {
          img: '项目规模',
          title: '房源项目规模'
        },
        {
          img: '运营效率',
          title: '运营效率'
        },
        {
          img: '产业生态',
          title: '产业生态'
        }
      ],
      // 房源项目规模
      projectScaleData: [
        {
          leftImg: '累计房源',
          projectTitle: '累计房源数据',
          projectNum: 6
        },
        {
          leftImg: '持有房源',
          projectTitle: '当前持有房源数',
          projectNum: 5
        }, {
          leftImg: '运营房源',
          projectTitle: '可运营房源数',
          projectNum: 8
        }
      ],
      // 产业生态
      productEcology: [
        {
          leftImg: '商户管理',
          projectTitle: '合作运营商',
          projectNum: 6000
        },
        {
          leftImg: '智能设备',
          projectTitle: '接入智能设备数',
          projectNum: 50
        }
      ],
      selectStatus: false,
      isActive: 0,
      hide: true

    }
  },
  methods: {
    shiftImg(type, url) {
      if (type === 'right') {
        return new URL(`../assets/images/regionRight/${url}.png`, import.meta.url).href
      } else if (type === 'left') {
        return new URL(`../assets/images/regionLeft/${url}.png`, import.meta.url).href
      } else if (type === 'ecology') {
        return new URL(`../assets/images/productEcology/${url}.png`, import.meta.url).href
      }
    },
    activeClick(i) {
      this.isActive = i
    }
  },
  watch: {
    '$store.state.region.aroundHide'(val) {
      if (val) {
        this.isActive = 0
      }
    },
    '$store.state.region.selectName'(n, o) {
      if (n !== o && n !== 'china') {
        this.hide = false;
        setTimeout(() => {
          this.hide = true
        })
      }
    }
  }
}
</script>
<style scoped lang='less'>
.box_prefecture {
  position: absolute;
  width: 100%;
  left: 0;
  top: 0;

  .main_left {
    position: absolute;
    left: 0.66rem;
    top: 0.74rem;
    animation-duration: 1s;

    .left_item {
      width: 100%;
      height: 1.45rem;
      display: flex;
      flex-direction: column;
      align-items: center;
      margin-bottom: 0.4rem;
      cursor: pointer;
    }

    .pic {
      width: .975rem;
      height: 1.125rem;
      margin-bottom: 0.14rem;
    }

    .title {
      font-size: 0.18rem;
      font-weight: 400;
      line-height: 1;
      color: #cfcfcf;
      font-family: none;
    }

    .active {
      color: #00ffff;
    }
  }

  .main_right {
    position: absolute;
    right: 0.66rem;
    top: 0;
    animation-duration: 1s;
    // border: 1px solid lightcoral;

    .total_data {
      position: relative;
      background: url('../assets/images/regionRight/数据框.png') no-repeat;
      background-size: 100% 100%;
      height: 1.5rem;
      width: 3.75rem;
      display: flex;
      justify-content: space-between;
      padding: 0.40rem 0.29rem 0;
      margin-bottom: 0.85rem;

      .td_left {
        display: flex;
        flex-direction: column;
        align-items: center;


        .l_top {
          width: 0.4rem;
          height: 0.48rem;
          margin-bottom: 0.12rem;
        }

        .l_bottom {
          font-size: 0.23rem;
          line-height: 1.2;
          font-weight: 400;
          font-family: 'lt';
        }
      }

      .td_right {
        font-size: 0.9rem;
        line-height: 1;
        font-weight: 400;
        color: #FFF100;

        span {
          padding-left: .1rem;
          font-size: 0.4rem;
          font-family: none;
          color: #fff;
        }
      }
    }

    .project_scale {
      border-top: 0.0125rem solid #183258;
      width: 3.7rem;
      height: 5.46rem;

      .ps_item {
        display: flex;
        width: 100%;
        height: 1.81rem;
        padding: 0.45rem 0 0 0.16rem;
        border-bottom: 0.0125rem solid #183258;

        .project_left {
          width: 0.91rem;
          height: 0.91rem;
        }

        .project_right {
          padding-left: 0.23rem;
          font-style: italic;

          .r_title {
            font-size: 0.22rem;
            font-weight: 400;
            line-height: 1;
            margin-bottom: 0.25rem;
          }

          .r_num {
            font-size: 0.45rem;
            font-weight: 400;
            line-height: 1;
            color: #00DEFF;

            span {
              color: #fff;
              padding-left: .1rem;
              font-size: 0.3rem;
              font-family: none;
            }
          }
        }
      }
    }
  }

  .left3 {
    top: 2rem;
  }
}
</style>