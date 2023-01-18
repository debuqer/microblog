<template>
  <div class="timeline-entry">
    <div class="timeline-stat">
      <div class="timeline-icon"><img :src="item.profile" alt="Profile picture">
      </div>
      <div class="timeline-time"><a href="#" class="btn-link text-semibold">{{ item.author }}</a><br/>{{ item.time }}</div>
    </div>
    <div class="timeline-label">
      <p class="mar-no pad-btm"> {{ item.description ?? 'commented on' }} when <a href="#">{{ item.label }}</a>
      </p>
      <div v-if="item.image" >
        <img class="img-fluid" :src="item.image">
      </div>
      <blockquote class="bq-sm bq-open">{{ item.message }}</blockquote>
      <div>
        <div class="activity__list__footer">
          <a href="#"> <i class="fa fa-thumbs-up"></i>{{ item.likes }}</a>
          <a href="#"> <i class="fa fa-comments"></i>{{ item.comments }}</a>
        </div>
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

<style>
.activity__list__footer {
    display: -ms-flexbox;
    display: -webkit-box;
    display: flex;
    margin-top: 23px;
    margin-left: 43px;
    padding: 13px 8px 0;
    color: #999;
    border-top: 1px dotted #ccc;
}

.activity__list__footer a {
    color: inherit;
}

.activity__list__footer a + a {
    margin-left: 15px;
}

.activity__list__footer i {
    margin-right: 8px;
}

.activity__list__footer a:hover {
    color: #222;
}

.activity__list__footer span {
    margin-left: auto;
}
</style>
