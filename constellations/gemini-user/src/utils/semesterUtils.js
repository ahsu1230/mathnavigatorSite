export const sortedSemesterInsert = (array, value) => {
    let low = 0;
    let high = array.length;

    while (low < high) {
        let mid = (low + high) >>> 1;
        if (compareSemesters(array[mid].semesterId, value.semesterId))
            low = mid + 1;
        else high = mid;
    }
    array.splice(low, 0, value);
    return array;
};

const compareSemesters = (semester1, semester2) => {
    const seasonMap = {
        spring: 0,
        summer: 1,
        fall: 2,
        winter: 3,
    };
    let [year1, season1] = semester1.split("_");
    let [year2, season2] = semester2.split("_");
    season1 = seasonMap[season1];
    season2 = seasonMap[season2];

    if (year1 == year2 && season1 < season2) return true;
    return year1 < year2;
};
