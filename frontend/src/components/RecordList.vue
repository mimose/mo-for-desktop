<template>
  <v-container fluid>
    <div
      v-if="recordGroup != null && recordGroup.length > 0"
      class="pl-14 pr-6"
      style="position: absolute;height: 90%; width:100%; overflow: auto; z-index: 1"
    >

      <v-tabs v-model="tabModel" centered height=50 fixed-tabs slider-color="accent darken-1">
        <v-tab
          v-for="group in recordGroup"
          :key="group.groupKey"
          :href="`#tab-${group.groupKey}`"
          @click="triggerExpand(group.groupKey)"
        >
          {{ group.groupName }}


          <v-progress-circular
            :value="progress(group.groupKey)"
            class="ml-4"
            size="20"
            width="10"
          ></v-progress-circular>
        </v-tab>
      </v-tabs>

      <v-tabs-items style="background: none" v-model="tabModel">
        <v-tab-item
          v-for="group in recordGroup"
          :key="group.groupKey"
          :value="`tab-${group.groupKey}`"
        >
          <!-- <v-card flat> -->
            <!-- <v-card-actions> -->
              <v-expansion-panels
                v-if="group.groupValues.length > 0"
                accordion
                multiple
                focusable
                :value="expands"
              >
                <template v-for="(record, i) in group.groupValues">
                  <v-expansion-panel
                    :key="`${i}-${record.key}`"
                    :readonly="record.readonly"
                  >
                    <v-expansion-panel-header disable-icon-rotate style="padding: 5px 10px 0px 10px">
                      <v-checkbox
                        v-model="record.done"
                        @change="check(record)"
                        :color="(record.done && 'grey') || 'primary'"
                        style="max-width: 5%"
                      ></v-checkbox>
                      <span
                        :class="
                          (record.done && 'grey--text') || 'primary--text'
                        "
                        class="ml-4"
                        v-text="record.title"
                      ></span>
                      <template v-slot:actions>
                        <v-icon dense v-if="record.done" color="success">
                          mdi-check
                        </v-icon>
                        <v-icon dense v-if="!record.done" color="primary"></v-icon>
                      </template>
                    </v-expansion-panel-header>
                    <v-expansion-panel-content eager>
                      <v-card elevation="0" flat >
                        <v-card-text class="pb-0">
                          <div v-if="record.noticeTime"><span>notice me at</span>&nbsp;&nbsp;{{record.noticeTime}}</div>
                        </v-card-text>
                        <v-card-text class="pt-0 pb-0 text--primary" v-text="record.content"></v-card-text>
                        <v-card-actions class="pt-0 pb-0">
                          <v-btn
                            icon
                            plain
                            right
                            absolute
                            color="teal accent-4"
                            @click="openEditRecord(record)"
                          >
                          <v-icon dark small>
                            mdi-pencil
                          </v-icon>
                          </v-btn>

                          <v-btn
                            plain
                            right
                            fixed
                            icon
                            color="teal accent-4"
                            @click="openDelRecord(record)"
                          >
                          <v-icon dark small>
                            mdi-minus
                          </v-icon>
                          </v-btn>
                        </v-card-actions>
                      </v-card>
                    </v-expansion-panel-content>
                  </v-expansion-panel>
                </template>
              </v-expansion-panels>
            <!-- </v-card-actions> -->
          <!-- </v-card> -->
        </v-tab-item>
      </v-tabs-items>

      <!--  -->
      <v-dialog
        v-model="delWindow"
        persistent
        max-width="290"
      >
        <v-card>
          <v-card-title class="text-h5">
            删除任务?
          </v-card-title>
          <v-card-text>
            <v-alert
                dense
                outlined
                type="error"
                :value="delErrorMessageShow">
              {{delErrorMessage}}
              </v-alert>
              任务删除后将无法恢复，是否删除？
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              :loading="delLoading"
              color="green darken-1"
              text
              @click="delWindow = false"
            >
              取消
            </v-btn>
            <v-btn
              :loading="delLoading"
              color="green darken-1"
              text
              @click="deleteRecord"
            >
              确定
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>

    <div
      class="pl-14 pr-6"
      style="bottom: 5px; position: absolute; width: 100%; z-index: 2"
    >
      <v-bottom-sheet v-model="editWindow" persistent eager>
        <template v-slot:activator="{ on, attrs }">
          <!-- <v-btn width="100%" color="accent darken-1" v-bind="attrs" v-on="on">
            新建任务
          </v-btn> -->
          <div class="text-center">
            <v-btn
              width="100%"
              class="my-3"
              color="accent darken-2"
              v-bind="attrs" v-on="on"
              @click="openCreateRecord"
              text
            >
              <v-icon dark>
                mdi-plus
              </v-icon>
              新建任务
            </v-btn>
          </div>
        </template>
        <v-sheet class="text-center" min-height="420px">
          <v-card 
            :loading="editLoading?`primary`:`none`"
            class="mx-auto"
            elevation="0" 
            flat
          >
          <v-form
            ref="editWindowRef">
            <v-card-text class="pl-14 pr-14">
              <v-alert
                dense
                outlined
                type="error"
                :value="editErrorMessageShow">
              {{editErrorMessage}}
              </v-alert>
              <v-text-field 
                :disabled="editLoading" 
                v-model="editRecord.title" 
                counter
                label="主题" 
                :rules="titleRules" 
                hide-details="auto"></v-text-field>
              <v-textarea 
                :disabled="editLoading" 
                v-model="editRecord.content" 
                counter
                no-resize
                rows="1"
                :rules="contentRules"
                label="详细内容" 
                hide-details="auto"></v-textarea>
              <v-dialog
                ref="noticeDateDialog"
                v-model="pickDateModal"
                :return-value.sync="editRecord.noticeDate"
                persistent>
                <template v-slot:activator="{ on, attrs }">
                  <v-text-field
                    :disabled="editLoading"
                    v-model="editRecord.noticeDate"
                    label="提示日期"
                    append-icon="mdi-calendar"
                    readonly
                    v-bind="attrs"
                    v-on="on"
                  ></v-text-field>
                </template>
                <v-date-picker
                  v-model="editRecord.noticeDate"
                  :show-current="editRecord.noticeDate === ''"
                  type="date"
                  scrollable>
                  <v-spacer></v-spacer>
                  <v-btn
                    text
                    color="primary"
                    @click="pickDateModal = false;editRecord.noticeDate='';$refs.noticeDateDialog.save(editRecord.noticeDate)">
                    清除
                  </v-btn>
                  <v-btn
                    text
                    color="primary"
                    @click="$refs.noticeDateDialog.save(editRecord.noticeDate)">
                    保存
                  </v-btn>
                </v-date-picker>
              </v-dialog>

              <v-dialog
                ref="noticeDateTimeDialog"
                v-model="pickDateTimeModal"
                :return-value.sync="editRecord.noticeDateTime"
                persistent>
                <template v-slot:activator="{ on, attrs }">
                  <v-text-field
                    :disabled="editLoading"
                    v-model="editRecord.noticeDateTime"
                    label="提示时间"
                    append-icon="mdi-clock"
                    readonly
                    v-bind="attrs"
                    v-on="on"
                  ></v-text-field>
                </template>
                <v-time-picker
                  v-model="editRecord.noticeDateTime"
                  :show-current="editRecord.noticeDateTime === ''"
                  use-seconds
                  format="24hr"
                  scrollable>
                  <v-spacer></v-spacer>
                  <v-btn
                    text
                    color="primary"
                    @click="pickDateTimeModal = false;editRecord.noticeDateTime='';$refs.noticeDateTimeDialog.save(editRecord.noticeDateTime)">
                    清除
                  </v-btn>
                  <v-btn
                    text
                    color="primary"
                    @click="$refs.noticeDateTimeDialog.save(editRecord.noticeDateTime)">
                    保存
                  </v-btn>
                </v-time-picker>
              </v-dialog>
            </v-card-text>
          </v-form>
          </v-card>
          <!-- <v-divider class="mx-4"></v-divider> -->
          <v-btn :disabled="editLoading" class="mt-1 pl-14" text left fixed color="primary" @click="saveRecord">
            save
          </v-btn>
          <v-btn :disabled="editLoading" class="mt-1 pr-14" text right fixed color="accent" @click="editWindow = !editWindow">
            close
          </v-btn>
        </v-sheet>
      </v-bottom-sheet>
    </div>
  </v-container>
</template>

<script>
// const notifier = require('node-notifier')
export default {
  data: () => ({
    tabModel: "tab-notice",
    expands: [],
    recordGroup: [
      {
        groupKey: "notice",
        groupName: "notice",
        groupValues: [],
      },
    ],
    delWindow: false,
    delLoading: false,
    delErrorMessageShow: false,
    delErrorMessage: "",
    delRecord: {},

    editLoading: false,
    editErrorMessageShow: false,
    editErrorMessage: "",
    editWindow: false,
    editRecord: {},
    pickDateModal: false,
    pickDateTimeModal: false,
    titleRules: [
      value => !!value || '请输入主题信息',
      value => value.length <= 20 || '超过字数限制'
    ],
    contentRules: [
      value => value.length <= 35 || '超过字数限制'
    ]
  }),
  computed: {
    completedRecordCount() {
      return function(groupKey) {
        if(groupKey == 'notice') {
          return this.recordGroup[0].groupValues.filter((task) => task.done).length;
        } else {
          return this.recordGroup[1].groupValues.filter((task) => task.done).length;
        }
      }
    },
    recordCount() {
      return function(groupKey) {
        if(groupKey == 'notice') {
          return this.recordGroup[0].groupValues.length;
        } else {
          return this.recordGroup[1].groupValues.length;
        }
      }
    },
    remainingRecordCount() {
      return function(groupKey) {
        return this.recordCount(groupKey) - this.completedRecordCount(groupKey);
      }
    },
    progress() {
      return function(groupKey) {
        return (this.completedRecordCount(groupKey) / this.recordCount(groupKey)) * 100;
      }
    },
  },
  mounted: function() {
    this.initWindow()
  },
  methods: {
    async initWindow() {
      let vm = this;
      // 进行全局遮罩，在查询时进行
      vm.$emit('update:appOverlaySync', true)
      await vm.completeRecordGroup('notice', 0);
      await vm.completeRecordGroup('note', 1);
      vm.$emit('update:appOverlaySync', false)
    },
    async completeRecordGroup(groupKey, recordKey) {
      let vm = this;
      let noticeRecordGroup = vm.recordGroup.find(item => item.groupKey === groupKey);
      if(noticeRecordGroup == null) {
        return;
      }
      noticeRecordGroup.groupValues = [];
      if(vm.selectedSpace === '' || vm.selectedSpace === null) {
        return;
      }
      let noticesResp = await window.backend.Mo.ListRecord(recordKey);
      if(noticesResp.code === 200) {
        if(noticesResp.data != null) {
          noticesResp.data.forEach(element => {
            console.log(element)
            if(element.noticeTime !== "" && element.noticeTime.startsWith("0001")) {
              element.noticeTime = "";
            }
            let value = {
              key: element.key,
              done: element.done,
              title: element.title,
              content: element.content,
              noticeTime: element.noticeTime,
              readonly: element.content == null || element.content == ''
            }
            noticeRecordGroup.groupValues.push(value);
          });
        }
      }
    },
    async check(record) {
      // let vm = this;
      let resp;
      if(record.done) {
        // 设置为完成
        resp = await window.backend.Mo.DoneRecord(record.key);
      } else {
        resp = await window.backend.Mo.UndoneRecord(record.key);
      }
      if(resp.code !== 200) {
        record.done = !record.done;
      }
    },
    openDelRecord(record) {
      let vm = this;
      vm.delErrorMessageShow = false;
      vm.delErrorMessage = "";
      vm.delRecord = {};
      vm.delRecord = record;
      vm.delWindow = true;
    },
    async deleteRecord() {
      let vm = this;
      vm.delLoading = true;
      vm.delErrorMessageShow = false;
      vm.delErrorMessage = "";

      let delResp = await window.backend.Mo.RemoveRecord(vm.delRecord.key);
      if(delResp.code === 200) {
        vm.initWindow();
        vm.delRecord = {};
        vm.delWindow = false;
      } else {
        vm.delErrorMessage = delResp.desc;
        vm.delErrorMessageShow = true;
      }
      vm.delLoading = false;
    },
    openEditRecord(record) {
      let vm = this;
      vm.editErrorMessageShow = false;
      vm.editErrorMessage = "";
      vm.editRecord = {
        key: record.key,
        title: record.title,
        content: record.content,
        noticeDate: vm.splitTimeToGetDate(record.noticeTime),
        noticeDateTime: vm.splitTimeToGetDateTime(record.noticeTime)
      };
      vm.editWindow = !vm.editWindow;
    },
    openCreateRecord() {
      let vm = this;
      vm.editRecord = {};
      vm.editErrorMessageShow = false;
      vm.editErrorMessage = "";
      vm.editWindow = !vm.editWindow;
    },
    async saveRecord() {
      let vm = this;

      if(!vm.$refs.editWindowRef.validate()) {
        vm.editErrorMessageShow = true;
        vm.editErrorMessage = "请根据提示完善信息";
        return
      }

      vm.editLoading = true;
      vm.editErrorMessageShow = false;
      vm.editErrorMessage = "";

      if(vm.editRecord.noticeDate !== "" && vm.editRecord.noticeDateTime === "") {
        vm.editErrorMessageShow = true;
        vm.editErrorMessage = "请补充完整提示时间";
      }

      if(vm.editRecord.noticeDate === "" && vm.editRecord.noticeDateTime !== "") {
        vm.editErrorMessageShow = true;
        vm.editErrorMessage = "请补充完整提示日期";
      }

      if(vm.editRecord.noticeDate !== "") {
        vm.editRecord.noticeTime = vm.editRecord.noticeDate + " " + vm.editRecord.noticeDateTime
      }
      let saveResp = await window.backend.Mo.NewOrUpdateRecord(JSON.stringify(vm.editRecord));
      if(saveResp.code === 200) {
        vm.initWindow();
        vm.editRecord = {};
        vm.editWindow = !vm.editWindow
      } else {
        vm.editErrorMessage = saveResp.desc;
        vm.editErrorMessageShow = true;
      }
      vm.editLoading = false;
    },
    // createRecord() {
    //   let vm = this;
    //   // 进行全局遮罩，在查询时进行
    //   vm.$emit('update:appOverlaySync', true)

    //   this.recordGroup[0].groupValues.push({
    //     done: false,
    //     text: this.newNoticeTask,
    //   });

    //   this.newNoticeTask = null;
    // },
    triggerExpand(groupKey) {
      let vm = this;
      if(vm.expands && vm.expands.length > 0) {
        vm.expands = []
      } else {
        let noticeRecordGroup = vm.recordGroup.find(item => item.groupKey === groupKey);
        let length = noticeRecordGroup.groupValues.length;
        let allIndex = [];
        if(length>0) {
          for (let index = 0; index < length; index++) {
            allIndex.push(index);
          }
        }
        vm.expands = allIndex
      }
    },
    splitTimeToGetDate(time) {
      if(time && time != null && time != "") {
        return time.substr(0, 10)
      }
      return ""
    },
    splitTimeToGetDateTime(time) {
      if(time && time != null && time != "") {
        return time.substr(11, time.length)
      }
      return ""
    },
    // testNotice() {
      // debugger;
      // String
      // notifier.notify('Message');
      // Object
      // notifier.notify({
      //   'title': 'My notification',
      //   'message': 'Hello, there!'
      // });

    // }
  },
};
</script>