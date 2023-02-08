<template>
  <div class="test-box" :style="{ '--size': size, '--slotSize': slotSize }" v-if="hiden">
    <div class="test">{{ animatedNumber }}</div>
    <slot></slot>
  </div>
</template>

<script>
import gsap from 'gsap'
export default {
  props: {
    value: {
      type: [String, Number],
      default: 0
    },
    size: {
      type: String,
    },
    slotSize: {
      type: String,
      // default: '0rem'
    },
    bigNum: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      num: 0,//初始值
      hiden: true,
    }
  },
  mounted() {
    this.animateText()
  }
  ,
  computed: {
    animatedNumber() {

      if (this.bigNum) {
        return this.num.toFixed(0).toString().replace(/(\d)(?=(\d{2})(?!\d))/g, '$1.');
      } else {
        return this.num.toFixed(0)
      }
      // return this.num.toFixed(0).toString().replace(/(\d)(?=(\d{2})(?!\d))/g, '$1.');

      // return this.largeValue(this.num.toFixed(0))
    }
  }
  ,
  watch: {
    value() {
      this.animateText()
    },
    '$store.state.region.selectName'(val) {
      this.hiden = false;
      setTimeout(() => {
        this.hiden = true
      })
    },
  },
  methods: {
    animateText() {
      gsap.to(this, {
        duration: 1.5,
        num: this.value,
        repeat: -1,
        repeatDelay: 6
      })

    },
    formatNumber(num) {
      return num.toString().replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1,');
    },
    // largeValue(val) {
    //   // 10W
    //   if (val > 100000 && val <= 1000000) {
    //     let num1 = val.toString().slice(0, 2);
    //     let num2 = val.toString().slice(2, 4);
    //     return `${num1}.${num2}`
    //   } else if (val > 1000000) {
    //     let num1 = val.toString().slice(0, 3);
    //     let num2 = val.toString().slice(3, 5);
    //     return `${num1}.${num2}`
    //   } else {
    //     return val
    //   }
    // }
  }
}
</script>

<style lang="less" scoped>
.test-box {
  display: flex;
  align-items: flex-end;
  justify-content: flex-start;

  span {
    line-height: 1.3;
    font-size: var(--slotSize, 0.1rem);
    padding-left: .0375rem;
  }

  .test {
    font-size: var(--size, 0.3rem);
    line-height: 1;
    font-weight: 400;
  }
}
</style>