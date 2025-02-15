import request from "@/request/index";

// params: { ownerid: 1, quesid: 1, quesbankid: 1 }
export function getQuestions(params) {
  return request({
    url: "/question/get",
    method: "get",
    params,
  });
}
/* data: { subject: "物理", class: "高一", difficulty: "中等", choice_num: 1, calc_num: 1, short_ans_num: 1, knowledge: "牛顿第二定律" } */
export function genQuestions(data) {
  return request({
    url: "/question/gen",
    method: "post",
    data,
  });
}
// data: {quesid,answer}
export function checkAnswer(data) {
  return request({
    url: "/question/check",
    method: "post",
    data,
  });
}
