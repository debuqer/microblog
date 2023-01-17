<template>
  <div class="panel">
    <div class="panel-body">
      <textarea class="form-control" rows="2" placeholder="What are you thinking?"></textarea>
      <div class="mar-top clearfix">
        <button class="btn btn-sm btn-primary pull-right" type="submit"><i class="fa fa-pencil fa-fw"></i>
          Share</button>
        <a class="btn btn-trans btn-icon fa fa-video-camera add-tooltip" href="#" data-original-title="Add Video"
          data-toggle="tooltip"></a>
        <a class="btn btn-trans btn-icon fa fa-camera add-tooltip" href="#" data-original-title="Add Photo"
          data-toggle="tooltip"></a>
        <a class="btn btn-trans btn-icon fa fa-file add-tooltip" href="#" data-original-title="Add File"
          data-toggle="tooltip"></a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PostSomething',
  props: {
  }
}
</script>
