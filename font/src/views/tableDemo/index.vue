<template>
    <div class='system-user-container'>
        <app-table ref='appTableRef' tableName="adminTable" :config='tableConfig'>
            <template #header-action='{ loading }'>
                <el-button
                    :loading='loading'
                    type='primary'
                    icon='el-icon-plus'
                    size='small'
                    @click='handleAdd'
                >
                    新增
                </el-button>
                <el-button
                    :loading='loading'
                    type='danger'
                    icon='el-icon-delete'
                    size='small'
                    @click='handleDelete'
                >
                    删除
                </el-button>
            </template>
            <template #is-active>
                <el-table-column
                    :width='100'
                    label='用户状态'
                    align='center'
                >
                    <template #default='scope'>
                        <el-tag v-if='scope.row.isActive' type='success'>正常</el-tag>
                        <el-tag v-else type='danger'>禁用</el-tag>
                    </template>
                </el-table-column>
            </template>
        </app-table>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, reactive } from 'vue'
import { ElNotification, ElMessageBox, ElMessage } from 'element-plus'

import AppTable from '@/components/AppTable.vue'
import { getTableList, deleteTableItem } from '@/api/components/index'

export default defineComponent({
    name: 'SystemUser',
    components: {
        AppTable
    },
    setup() {
        const tableConfig = reactive({
            tableColumn: [
        
                {
                    label: '序号',
                    type: 'index',
                    width: 50,
                    align: 'center'
                },
                {
                    align: 'center',
                    label: 'id',
                    prop: 'id',
                    width: 50,
                    sortable: true
                },
                {
                    align: 'center',
                    label: '路由',
                    prop: 'route',
                    width: 500,
                    sortable: true
                },
                {
                    slot:'is-active'
                },
                {
                    width: 120,
                    align: 'center',
                    fixed:'right',
                    label: '操作',
                    action: ['editBox', 'delete']
                }
            ],
            tableSearchBtnName: {
                reset: false
            },
            tableListApi: params => getTableList(params),
            tableListParams: {},
            tableDeleteApi: (id: number) => deleteTableItem(id),
            tableDeleteParams: {},
            tableSearch: [
                {
                    label: '路径',
                    type: 'text',
                    key: 'route',
                    inputType: 'text',
                    placeholder: '请输入登录账号',
                    clearable: true
                }
                // {
                //   label: '用户状态',
                //   type: 'select',
                //   key: 'isActive',
                //   value: 3,
                //   placeholder: '请选择用户状态',
                //   options: [
                //     {
                //       label: '正常',
                //       value: 1
                //     },
                //     {
                //       label: '禁用',
                //       value: 0
                //     },
                //     {
                //       label: '全部',
                //       value: 3
                //     }
                //   ]
                // }
            ]
        })
        const appTableRef = ref(null)

        const handleAdd = () => {
            const { selectionRow } = appTableRef.value
            if (selectionRow.length === 0) {
                ElNotification({
                    title: '操作失败',
                    message: '请选择数据再进行操作！',
                    type: 'warning'
                })
                return false
            }
        }
        const handleDelete = () => {
            const { selectionRow } = appTableRef.value
            if (selectionRow.length === 0) {
                ElNotification({
                    title: '操作失败',
                    message: '请选择数据再进行操作！',
                    type: 'warning'
                })
                return false
            }
            ElMessageBox.confirm('此操作将会永久删除数据，是否继续？', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                request.delete('/user/list', {
                    data: {
                        list: selectionRow.map(i => i.id)
                    }
                }).then(() => {
                    appTableRef.value.onSearchSubmit()
                })
            })
                .catch(() => {
                    ElMessage({
                        type: 'info',
                        message: '取消操作！'
                    })
                })
        }

        return {
            tableConfig,
            appTableRef,
            handleAdd,
            handleDelete
        }
    }
})
</script>

<style lang="scss" scoped>
.system-user-container {
  height: 100%;
}
</style>
