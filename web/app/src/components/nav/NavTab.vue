<template>
    <div class="card card-primary card-outline card-outline-tabs mb-0">
        <div class="card-header p-0 pt-1 border-bottom-0">
            <ul class="nav nav-tabs">
                <li v-for="nav in navs" :key="nav" class="nav-item">
                    <router-link class="nav-link" v-bind:class="{ active: nav.isActive }" :to="{ name: nav.link}">{{ nav.name }}</router-link>
                </li>
            </ul>  
        </div>
    </div>
</template>

<script>
export default {
    props: {
        navs: {
            type: Array,
            required: true
        }
    }
}
</script>

<style>

</style>