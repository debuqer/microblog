<template>
  <div class="row">
    <div class="col col-md-4">
      <select class="form-control" v-model="defaultArea" @change="reset">
        <option>AD</option>
        <option>BC</option>
      </select>
    </div>
    <div class="col col-md-8">
      <input type="text" class="form-control mr-2" v-model="defaultDate" v-maska :data-maska="format" :placeholder="placeholder" v-validate="{ date_format: 'dd/MM/yyyy', date_between:['01/01/1990',maxStartDate] }" />
    </div>
  </div>
</template>

<script setup>
import { vMaska } from 'maska'
</script>

<script>
export default {
  data: function () {
    return {
      defaultArea: 'AD',
      defaultDate: ''
    }
  },
  computed: {
    format: function () {
      return this.area === 'AD' ? '####/##/##' : '####'
    },
    placeholder: function () {
      return this.area === 'AD' ? 'E.g. 1998/11/24' : 'E.g. 2500'
    }
  },
  methods: {
    reset: function () {
      this.defaultDate = ''
    }
  },
  name: 'SuperDatePicker',
  props: {
    area: String,
    date: String
  }
}
</script>

<style>
/* The switch - the box around the slider */
.switch {
  position: relative;
  display: inline-block;
  width: 60px;
  height: 34px;
}

/* Hide default HTML checkbox */
.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

/* The slider */
.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  -webkit-transition: .4s;
  transition: .4s;
}

.slider:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  -webkit-transition: .4s;
  transition: .4s;
}

input:checked + .slider {
  background-color: #2196F3;
}

input:focus + .slider {
  box-shadow: 0 0 1px #2196F3;
}

input:checked + .slider:before {
  -webkit-transform: translateX(26px);
  -ms-transform: translateX(26px);
  transform: translateX(26px);
}

/* Rounded sliders */
.slider.round {
  border-radius: 34px;
}

.slider.round:before {
  border-radius: 50%;
}
</style>
