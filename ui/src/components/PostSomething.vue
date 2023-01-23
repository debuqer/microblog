<template>
  <div class="panel">
    <div class="panel-body">
      <div class="row">
        <div class="col col-md-8">
          <input type="text" class="form-control" placeholder="TL/LR" v-model="label"/>
        </div>
        <div class="col col-md-4">
          <SuperDatePicker area="AD" :date="''"></SuperDatePicker>
        </div>
        <div class="col col-md-8 mt-2">
          <textarea class="form-control" rows="2" :placeholder="textBoxPhrase" v-model="message"></textarea>
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
    tags: Array
  },
  data: () => {
    return {
      format: 'yyyy/MM/dd',
      label: '',
      message: '',
      date: ''
    }
  },
  computed: {
    tagsCommaseperated: function () {
      return this.tags
    },
    dateOnText: function () {
      return this.date ?? 'many years ago'
    },
    textBoxPhrase: function () {
      if (this.tagsCommaseperated.length === 0) {
        return 'What happend for everything in the world?'
      }
      return 'What happend for ' + this.tagsCommaseperated + '?'
    }
  },
  methods: {
    writeHistory: function () {
      console.log(this.message, this.date, this.tags)

      this.message = ''
      this.date = ''
      this.label = ''
    }
  }
}
</script>
