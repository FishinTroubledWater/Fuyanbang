<template>
  <div class="home">
    <Navbar></Navbar>
  </div>
</template>

<script>
// @ is an alias to /src


import Navbar from "@/components/Navbar";

export default {
  name: 'HomeView',
  components: {
    Navbar

  }
}
</script>
<style>
.home{
  height: 100%;
}
</style>