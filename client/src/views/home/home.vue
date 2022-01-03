<template>
  <div>
    <Navbar></Navbar>
    <router-view></router-view>
  </div>
</template>

<script>
import Navbar from "@/components/common/navBar.vue"
export default {
  name:"home",
  components: {
    Navbar,
  }
}
</script>

<style scoped>

</style>