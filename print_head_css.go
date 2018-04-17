package main

const printHeadStyle = `<style type="text/css">.d2h-file-diff .d2h-ins.d2h-change {
    background-color: #ded;
}

.d2h-ins {
    background-color: #dfd;
    border-color: #b4e2b4;
}

.d2h-code-line {
    display: inline-block;
    white-space: nowrap;
    padding: 0 10px;
    margin-left: 80px;
}

.d2h-wrapper {
    text-align: left;
}

.d2h-file-diff {
    overflow-x: scroll;
    overflow-y: hidden;
}

.d2h-diff-table {
    width: 100%;
    border-collapse: collapse;
    font-family: Menlo, Consolas, monospace;
    font-size: 13px;
}

table {
    background-color: transparent;
}

tbody {
    display: table-row-group;
    vertical-align: middle;
    border-color: inherit;
}

tr {
    display: table-row;
    vertical-align: inherit;
    border-color: inherit;
}

.selecting-left .d2h-code-line,
.selecting-left .d2h-code-line *,
.selecting-left .d2h-code-side-line,
.selecting-left .d2h-code-side-line *,
.selecting-right td.d2h-code-linenumber,
.selecting-right td.d2h-code-linenumber *,
.selecting-right td.d2h-code-side-linenumber,
.selecting-right td.d2h-code-side-linenumber * {
    -webkit-touch-callout: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
}

.d2h-diff-tbody>tr>td {
    height: 20px;
    line-height: 20px;
}

.d2h-info {
    background-color: #f8fafd;
    color: rgba(0, 0, 0, .3);
    border-color: #d5e4f2;
}

.d2h-code-linenumber {
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    position: absolute;
    width: 86px;
    padding-left: 2px;
    padding-right: 2px;
    background-color: #fff;
    color: rgba(0, 0, 0, .3);
    text-align: right;
    border: solid #eee;
    border-width: 0 1px 0 1px;
    cursor: pointer;
}

.d2h-diff-tbody>tr>td {
    height: 20px;
    line-height: 20px;
}

.d2h-info {
    background-color: #f8fafd;
    color: rgba(0, 0, 0, .3);
    border-color: #d5e4f2;
}

td,
th {
    padding: 0;
}

* {
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
}

.d2h-code-line {
    display: inline-block;
    white-space: nowrap;
    padding: 0 10px;
    margin-left: 80px;
}

.d2h-code-line-prefix {
    display: inline;
    background: 0 0;
    padding: 0;
    word-wrap: normal;
    white-space: pre;
}

.d2h-code-line-ctn {
    display: inline;
    background: 0 0;
    padding: 0;
    word-wrap: normal;
    white-space: pre;
}

.d2h-file-diff .d2h-del.d2h-change {
    background-color: #fdf2d0;
}

.d2h-del {
    background-color: #fee8e9;
    border-color: #e9aeae;
}

.d2h-code-line {
    display: inline-block;
    white-space: nowrap;
    padding: 0 10px;
    margin-left: 80px;
}

.d2h-ins {
    background-color: #dfd;
    border-color: #b4e2b4;
}

.d2h-code-line del,
.d2h-code-side-line del {
    display: inline-block;
    margin-top: -1px;
    text-decoration: none;
    background-color: #ffb6ba;
    border-radius: .2em;
}

.d2h-code-line ins,
.d2h-code-side-line ins {
    display: inline-block;
    margin-top: -1px;
    text-decoration: none;
    background-color: #97f295;
    border-radius: .2em;
    text-align: left;
}
.d2h-file-header {
    padding: 5px 10px;
    border-bottom: 1px solid #d8d8d8;
    background-color: #f7f7f7;
}

.d2h-file-name-wrapper {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    width: 100%;
    font-family: "Source Sans Pro","Helvetica Neue",Helvetica,Arial,sans-serif;
    font-size: 15px;
}.d2h-icon-wrapper {
    line-height: 31px;
}
svg:not(:root) {
    overflow: hidden;
}

.d2h-icon {
    vertical-align: middle;
    margin-right: 10px;
    fill: currentColor;
}
.d2h-file-name {
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow-x: hidden;
    line-height: 21px;
}
.d2h-changed-tag {
    border: #d0b44c 1px solid;
}

.d2h-tag {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    font-size: 10px;
    margin-left: 5px;
    padding: 0 2px;
    background-color: #fff;
}
.d2h-changed {
    color: #d0b44c;
}
</style>`
