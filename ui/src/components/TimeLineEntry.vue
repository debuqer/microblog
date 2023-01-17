<template>
  <div class="timeline-entry">
    <div class="timeline-stat">
      <div class="timeline-icon"><img :src="item.profile" alt="Profile picture">
      </div>
      <div class="timeline-time">{{ item.time }}</div>
    </div>
    <div class="timeline-label">
      <p class="mar-no pad-btm"><a href="#" class="btn-link text-semibold">{{ item.author }}.</a> commented on <a href="#">{{ item.label }}</a>
      </p>
      <blockquote class="bq-sm bq-open">{{ item.message }}</blockquote>
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
