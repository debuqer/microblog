<template>
  <div class="timeline-entry">
    <div class="timeline-stat">
      <div class="timeline-icon"><img src="https://bootdey.com/img/Content/avatar/avatar6.png" alt="Profile picture">
      </div>
      <div class="timeline-time">3 Hours ago</div>
    </div>
    <div class="timeline-label">
      <p class="mar-no pad-btm"><a href="#" class="btn-link text-semibold">Lisa D.</a> commented on <a href="#">The
          Article</a>
      </p>
      <blockquote class="bq-sm bq-open">Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy
        nibh euismod tincidunt.</blockquote>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TimeLineEntry',
  props: {
    msg: String
  }
}
</script>
