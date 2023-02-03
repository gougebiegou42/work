<template>
  <div class="select_box">
    <el-select v-model="value" placeholder="全国" class="select" :popper-append-to-body="false" @change="selectHit">
      <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
      </el-option>
    </el-select>
  </div>
</template>

<script>
export default {
  data() {
    return {
      options: [
        {
          value: 'china',
          label: '全国'
        }, {
          value: '深圳市',
          label: '深圳市'
        }, {
          value: '广州市',
          label: '广州市'
        }, {
          value: '上海市',
          label: '上海市'
        }, {
          value: '北京市',
          label: '北京市'
        }],
      value: ''
    }
  },
  mounted() {

  },
  methods: {
    selectHit(e) {
      if (e === 'china') {
        this.$store.commit("region/hideDisable", true);
        this.$store.commit("region/selectDisable", e)
      } else {
        this.$store.commit("region/hideDisable", false);
        this.$store.commit("region/selectDisable", e)
      }
    },
  },
  watch: {
    '$store.state.region.selectName'(val) {
      this.options.map((item) => {
        if (val === item.value) {
          return this.value = val
        }
      })
    }
  }
}
</script>
<style scoped lang='less'>
.select_box {
  position: absolute;
  top: 1.13rem;
  left: 0.41rem;
}

.select {
  background: url(../assets/images/区域选择/矩形.png) no-repeat;
  background-size: 100% 100%;
  width: 2.56rem;
  height: 0.45rem;
  line-height: 0.45rem;

  // padding-left: 0.21rem;
  /deep/.el-input__inner {
    background-color: transparent;
    height: .4rem;
    border: none;
    line-height: .2875rem;
    padding: 0 .2rem;
    color: #fff;
    font-size: .1625rem;

    &::placeholder {
      color: #fff;
    }
  }

  /deep/.el-input__icon {
    height: auto;
    line-height: .3125rem;
    color: #00A0E9;
  }

  // .select_icon {
  //   background: url(../assets/images/区域选择/下拉.png) no-repeat;
  //   background-size: 100% 100%;
  //   width: 0.18rem;
  //   height: 0.09rem;
  //   position: absolute;
  //   top: 0.18rem;
  //   right: 0.17rem;
  // }
}

/deep/.el-select-dropdown {
  border-color: #018ed4;
  background-color: #002561;
  left: 0 !important;


  .el-select-dropdown__item {
    font-family: none;
    color: #fff;
    font-size: .1625rem;
    font-weight: 0;
    padding: 0 .2rem;

    &:hover {
      background-color: #06346d;
    }

    &.hover {
      background-color: #06346d;
    }
  }

  .popper__arrow {
    background-color: #06346d;
    border-bottom-color: #018ed4;
    top: -0.1rem;

    &::after {
      // top: 2px;
      border-bottom-color: #06346d;
    }
  }
}
</style>