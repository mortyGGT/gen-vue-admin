import service from '@/utils/request'

// @Tags Annoucement
// @Summary 创建Annoucement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Annoucement true "创建Annoucement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /annouce/createAnnoucement [post]
export const createAnnoucement = (data) => {
  return service({
    url: '/annouce/createAnnoucement',
    method: 'post',
    data
  })
}

// @Tags Annoucement
// @Summary 删除Annoucement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Annoucement true "删除Annoucement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /annouce/deleteAnnoucement [delete]
export const deleteAnnoucement = (data) => {
  return service({
    url: '/annouce/deleteAnnoucement',
    method: 'delete',
    data
  })
}

// @Tags Annoucement
// @Summary 删除Annoucement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Annoucement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /annouce/deleteAnnoucement [delete]
export const deleteAnnoucementByIds = (data) => {
  return service({
    url: '/annouce/deleteAnnoucementByIds',
    method: 'delete',
    data
  })
}

// @Tags Annoucement
// @Summary 更新Annoucement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Annoucement true "更新Annoucement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /annouce/updateAnnoucement [put]
export const updateAnnoucement = (data) => {
  return service({
    url: '/annouce/updateAnnoucement',
    method: 'put',
    data
  })
}

// @Tags Annoucement
// @Summary 用id查询Annoucement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Annoucement true "用id查询Annoucement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /annouce/findAnnoucement [get]
export const findAnnoucement = (params) => {
  return service({
    url: '/annouce/findAnnoucement',
    method: 'get',
    params
  })
}

// @Tags Annoucement
// @Summary 分页获取Annoucement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Annoucement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /annouce/getAnnoucementList [get]
export const getAnnoucementList = (params) => {
  return service({
    url: '/annouce/getAnnoucementList',
    method: 'get',
    params
  })
}
