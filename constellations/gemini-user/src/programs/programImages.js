import srcMathMAA from "../../assets/card_images/mathMaa.jpg";
import srcMathCalculus from "../../assets/card_images/mathCalculus.jpg";
import srcMath1 from "../../assets/card_images/math1.jpg";
import srcMath2 from "../../assets/card_images/math2.jpg";
import srcMath3 from "../../assets/card_images/math3.jpg";

import srcEnglish1 from "../../assets/card_images/english1.jpg";
import srcEnglish2 from "../../assets/card_images/english2.jpg";
import srcEnglish3 from "../../assets/card_images/english3.jpg";
import srcEnglish4 from "../../assets/card_images/english4.jpg";
import srcEnglish5 from "../../assets/card_images/english5.jpg";

import srcCoding1 from "../../assets/card_images/coding1.jpg";
import srcCoding2 from "../../assets/card_images/coding2.jpg";
import srcCoding3 from "../../assets/card_images/coding3.jpg";
import srcCodingJava from "../../assets/card_images/codingJava.jpg";
import srcCodingWeb from "../../assets/card_images/codingWeb.jpg";

export const ImgSrcMap = {
    math: [srcMath1, srcMath2, srcMath3],
    english: [srcEnglish1, srcEnglish2, srcEnglish3, srcEnglish4, srcEnglish5],
    coding: [srcCoding1, srcCoding2, srcCoding3],
};

export const getImageForProgramClass = (programObj) => {
    const title = programObj.title;
    if (foundInString(title, "amc") || foundInString(title, "maa")) {
        return srcMathMAA;
    } else if (foundInString(title, "calc")) {
        return srcMathCalculus;
    } else if (foundInString(title, "java")) {
        return srcCodingJava;
    } else if (foundInString(title, "web")) {
        return srcCodingWeb;
    }
    return getRandomImageFrom(ImgSrcMap[programObj.subject]);
};

const getRandomImageFrom = (imgArray) => {
    const imgSrcIndex = Math.floor(Math.random() * imgArray.length);
    return imgArray[imgSrcIndex];
};

const foundInString = (string, substring) => {
    return string.toLowerCase().indexOf(substring.toLowerCase()) >= 0;
};
