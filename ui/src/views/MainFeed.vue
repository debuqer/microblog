<template>
  <link href="https://maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css" rel="stylesheet">
  <div class="container bootstrap snippets">
    <div class="row">
      <div class="col-md-3">
        <SideBar></SideBar>
      </div>
      <div class="col-md-7 col-md-offset-2">
        <TimeLine :items="items"></TimeLine>
      </div>
    </div>
  </div>
</template>

<script>
import TimeLine from '@/components/TimeLine.vue'
import SideBar from '@/components/SideBar.vue'
export default {
  name: 'MainFeed',
  components: {
    TimeLine,
    SideBar
  },
  data: () => {
    return {
      items: [
        {
          id: 1,
          profile: 'https://bootdey.com/img/Content/avatar/avatar6.png',
          label: 'Time was 22:53',
          author: 'Bagher',
          day: 'Yesterday',
          time: '1 Hour ago',
          description: 'USA attacked to China',
          message: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
          image: 'https://bootdey.com/img/Content/avatar/avatar6.png',
          likes: 0,
          comments: 12
        },
        {
          id: 2,
          profile: 'https://bootdey.com/img/Content/avatar/avatar6.png',
          label: 'Time was 22:53',
          author: 'Bagher',
          day: 'Today',
          time: '1 Hour ago',
          description: 'China Attacked to Spain',
          message: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
          likes: 2,
          comments: 12
        }
      ]
    }
  }
}
</script>
