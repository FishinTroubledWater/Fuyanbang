<template>
    <web-view :src="webUrl"></web-view>
</template>

<script>
    export default {
        data() {
            return {
                webUrl:''
            }
        },
        onLoad(val) {
            this.webUrl = val.url
        }
    }
</script>

<style>
</style>