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
          <TimeLine :items="items"></TimeLine>
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
  },
  data: () => {
    return {
      items: [
        {
          id: 1,
          profile: 'https://bootdey.com/img/Content/avatar/avatar6.png',
          label: '22:53',
          author: 'Bagher',
          time: '1 Hour ago',
          message: 'This is cold'
        },
        {
          id: 2,
          profile: 'https://bootdey.com/img/Content/avatar/avatar6.png',
          label: '22:53',
          author: 'Bagher',
          time: '1 Hour ago',
          message: 'This is cold'
        }
      ]
    }
  }
}
</script>
