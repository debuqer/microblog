<template>
  <div class="panel">
    <div class="panel-body">
      <div class="row">
        <div class="col col-md-8">
          <input type="text" class="form-control" placeholder="TL/LR" v-model="label"/>
        </div>
        <div class="col col-md-4">
          <SuperDatePicker area="AD" date=""></SuperDatePicker>
        </div>
        <div class="col col-md-8 mt-2">
          <textarea class="form-control" rows="2" placeholder="What happend?" v-model="message"></textarea>
        </div>
        <div class="col col-md-4 pt-4">
          <button class="btn btn-sm btn-primary pull-right" style="width:100%" type="submit" @click="writeHistory"><i class="fa fa-pencil fa-fw"></i>
              Write History</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import SuperDatePicker from '@/components/SuperDatePicker.vue'
export default {
  name: 'PostSomething',
  components: {
    SuperDatePicker
  },
  props: {
  },
  data: () => {
    return {
      format: 'yyyy/MM/dd',
      label: '',
      message: '',
      date: ''
    }
  },
  methods: {
    writeHistory: function () {
      console.log(this.message, this.date)

      this.message = ''
      this.date = ''
      this.label = ''
    }
  }
}
</script>
