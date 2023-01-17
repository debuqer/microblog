<template>
  <link href="https://maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css" rel="stylesheet">
  <div class="container bootstrap snippets bootdeys">
    <div class="col-md-7 col-md-offset-2">
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
      <div class="panel panel-body">
        <TimeLine></TimeLine>
      </div>
    </div>
  </div>
</template>

<script>
import TimeLine from '@/components/TimeLine.vue'
export default {
  name: 'MainFeed',
  components: {
    TimeLine
  }
}
</script>
