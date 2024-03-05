"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
(() => __awaiter(void 0, void 0, void 0, function* () {
    let iRootNode = {};
    let stSeparator = "\\";
    const gstPaths = [];
    const loadJson = () => {
        return new Promise((resolve, reject) => {
            const xhr = new XMLHttpRequest();
            xhr.open('GET', "dirLib.json", true);
            xhr.onreadystatechange = () => {
                if (4 != xhr.readyState) {
                    return;
                }
                if (200 != xhr.status) {
                    throw "Load Json Error";
                }
                iRootNode = JSON.parse(xhr.responseText);
                resolve("");
            };
            xhr.send(null);
        });
    };
    const createDir = (iNode, stParentDir) => {
        let path = iNode.text;
        if (stParentDir) {
            path = stParentDir + stSeparator + path;
        }
        gstPaths.push(path);
    };
    const parseNode = (iNode, stParentDir) => {
        if (iNode.text) {
            createDir(iNode, stParentDir);
        }
        if (stParentDir) {
            stParentDir += stSeparator;
        }
        iNode.text && (stParentDir += iNode.text);
        if (!iNode.children) {
            return;
        }
        for (const iChildNode of iNode.children) {
            parseNode(iChildNode, stParentDir);
        }
    };
    const generateBatFile = () => {
        let stResults = '';
        for (const path of gstPaths) {
            stResults += `md ${path}\r\n`;
        }
        const url = URL.createObjectURL(new Blob([stResults], { type: "text/plain" }));
        const domA = document.createElement('a');
        domA.setAttribute('href', url);
        domA.setAttribute('target', '_blank');
        domA.setAttribute('download', 'generate_dir.bat');
        document.body.appendChild(domA);
        domA.click();
        setTimeout(() => {
            document.body.removeChild(domA);
        }, 5000);
    };
    yield loadJson();
    parseNode(iRootNode, "");
    generateBatFile();
}))();
