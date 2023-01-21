<template>
  <div class="panel">
    <div class="panel-body">
      <div class="row">
        <div class="col col-md-10">
          <input type="text" class="form-control" placeholder="TL/LR" />
        </div>
        <div class="col col-md-2">
          <Datepicker v-model="date" text-input auto-apply close-on-auto-apply :enable-time-picker="false" :format="format"></Datepicker>
        </div>
        <div class="col col-md-10 mt-2">
          <textarea class="form-control" rows="2" placeholder="What happend?" v-model="message"></textarea>
        </div>
        <div class="col col-md-2 pt-4">
          <button class="btn btn-sm btn-primary pull-right" style="width:100%" type="submit" @click="writeHistory"><i class="fa fa-pencil fa-fw"></i>
              Write History</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PostSomething',
  props: {
  },
  data: () => {
    return {
      format: 'yyyy/MM/dd',
      message: '',
      date: ''
    }
  },
  methods: {
    writeHistory: function () {
      console.log(this.message, this.date)

      this.message = ''
      this.date = ''
    }
  }
}
</script>
