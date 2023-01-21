<template>
  <!-- begin timeline-time -->
  <div class="timeline-time">
    <span class="date">{{ item.day }}</span>
    <span class="time">{{ item.time }}</span>
  </div>
  <!-- end timeline-time -->
  <!-- begin timeline-icon -->
  <div class="timeline-icon">
    <a href="javascript:;">&nbsp;</a>
  </div>
  <!-- end timeline-icon -->
  <!-- begin timeline-body -->
  <div class="timeline-body">
    <div class="timeline-header">
        <span class="userimage"><img :src="item.profile" alt=""></span>
        <span class="username"><a href="javascript:;">{{ item.author }}</a> <small>{{ item.label }}</small></span>
        <span class="pull-right text-muted"></span>
    </div>
    <div class="timeline-content">
        <b>{{ item.description }}</b>
        <p>
          {{ item.message }}
        </p>
    </div>
    <div class="timeline-likes">
        <div class="stats-right">
          <span class="stats-text">{{ item.likes }} Likes</span>
          <span class="stats-text">{{ item.comments }} Comments</span>
        </div>
        <div class="stats">
          <span class="fa-stack fa-fw stats-icon">
          <i class="fa fa-circle fa-stack-2x text-danger"></i>
          <i class="fa fa-heart fa-stack-1x fa-inverse t-plus-1"></i>
          </span>
          <span class="fa-stack fa-fw stats-icon">
          <i class="fa fa-circle fa-stack-2x text-primary"></i>
          <i class="fa fa-thumbs-up fa-stack-1x fa-inverse"></i>
          </span>
          <span class="stats-total">4.3k</span>
        </div>
    </div>
    <div class="timeline-footer">
        <a href="javascript:;" class="m-r-15 text-inverse-lighter"><i class="fa fa-thumbs-up fa-fw fa-lg m-r-3"></i> Like</a>
        <a href="javascript:;" class="m-r-15 text-inverse-lighter"><i class="fa fa-comments fa-fw fa-lg m-r-3"></i> Comment</a>
    </div>
    <div class="timeline-comment-box">
        <div class="user"><img src="https://bootdey.com/img/Content/avatar/avatar6.png"></div>
        <div class="input">
          <form action="">
              <div class="input-group">
                <input type="text" class="form-control rounded-corner" placeholder="Write a comment...">
                <span class="input-group-btn p-l-10">
                <button class="btn btn-primary f-s-12 rounded-corner" type="button">Comment</button>
                </span>
              </div>
          </form>
        </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TimeLineEntry',
  props: {
    item: Object
  }
}
</script>
