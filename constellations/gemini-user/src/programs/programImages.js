import srcMathMAA from "../../assets/card_images/mathMaa.jpg";
import srcMathCalculus from "../../assets/card_images/mathCalculus.jpg";
import srcMath1 from "../../assets/card_images/math1.jpg";
import srcMath2 from "../../assets/card_images/math2.jpg";
import srcMath3 from "../../assets/card_images/math3.jpg";
import srcMath4 from "../../assets/card_images/math4.jpg";
import srcMath5 from "../../assets/card_images/math5.jpg";

import srcEnglishReading1 from "../../assets/card_images/englishReading1.jpg";
import srcEnglishReading2 from "../../assets/card_images/englishReading2.jpg";
import srcEnglishWriting1 from "../../assets/card_images/englishWriting1.jpg";
import srcEnglishWriting2 from "../../assets/card_images/englishWriting2.jpg";

import srcCoding1 from "../../assets/card_images/coding1.jpg";
import srcCoding2 from "../../assets/card_images/coding2.jpg";
import srcCoding3 from "../../assets/card_images/coding3.jpg";
import srcCodingJava from "../../assets/card_images/codingJava.jpg";
import srcCodingWeb from "../../assets/card_images/codingWeb.jpg";

const imagesEnglishReading = [srcEnglishReading1, srcEnglishReading2];
const imagesEnglishWriting = [srcEnglishWriting1, srcEnglishWriting2];
const imagesEnglishAll = imagesEnglishReading.concat(imagesEnglishWriting);

export const ImgSrcMap = {
    math: [srcMath1, srcMath2, srcMath3, srcMath4, srcMath5],
    english: imagesEnglishAll,
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
    } else if (
        foundInString(title, "writing") ||
        foundInString(title, "grammar") ||
        foundInString(title, "essay")
    ) {
        return getRandomImageFrom(imagesEnglishWriting);
    } else if (
        foundInString(title, "reading") ||
        foundInString(title, "comprehension")
    ) {
        return getRandomImageFrom(imagesEnglishReading);
    }
    return getRandomImageFrom(ImgSrcMap[programObj.subject]);
};

const getRandomImageFrom = (array) => {
    let imgArray = array || [];
    if (imgArray.length == 0) {
        return null;
    }
    const imgSrcIndex = Math.floor(Math.random() * imgArray.length);
    return imgArray[imgSrcIndex];
};

const foundInString = (string, substring) => {
    return (
        (string || "").toLowerCase().indexOf((substring || "").toLowerCase()) >=
        0
    );
};
