<template>
    <div class="content-header">
      <div class="container-fluid">
        <div class="row mb-2">
          <div class="col-sm-6">
            <h1 class="m-0 text-dark">{{title}}</h1>
          </div><!-- /.col -->
          <div class="col-sm-6">
            <ol class="breadcrumb float-sm-right">
              <li v-for="(each, key) in breadcrumbs" :key="key" class="breadcrumb-item" v-bind:class="{ active: each['isActive'] }">
                <a :href="each['url']">{{each["name"]}}</a>
              </li>
            </ol>
          </div><!-- /.col -->
        </div><!-- /.row -->
      </div><!-- /.container-fluid -->
    </div>  
</template>

<script>
export default {
  props : {
    breadcrumbs : {
      type : Array,
      default: () => [
        {"name" : "Home", "url" : "#", "isActive" : true}
      ]
    },
    title : {
      type : String,
      required : true
    },
  }
}
</script>

<style>

</style>