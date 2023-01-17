<template>
  <link href="https://maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css" rel="stylesheet">
  <div class="container bootstrap snippets bootdeys">
    <div class="row">
      <div class="col-md-3">
        <SideBar></SideBar>
      </div>
      <div class="col-md-7 col-md-offset-2">
        <PostSomething></PostSomething>
        <div class="panel panel-body">
          <TimeLine></TimeLine>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import PostSomething from '@/components/PostSomething.vue'
import TimeLine from '@/components/TimeLine.vue'
import SideBar from '@/components/SideBar.vue'
export default {
  name: 'MainFeed',
  components: {
    TimeLine,
    PostSomething,
    SideBar
  }
}
</script>
