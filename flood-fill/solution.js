/**
 * @param {number[][]} image
 * @param {number} sr
 * @param {number} sc
 * @param {number} newColor
 * @return {number[][]}
 */
var floodFill = function(image, sr, sc, newColor) {
    if(!image || !image.length || !image[sc].length) return;
    if(image[sr][sc] === newColor) return image;
    const sourceColor = image[sr][sc];
    fill(image, sr, sc, sourceColor, newColor);
    return image;
};

var fill = function(image, sr, sc, sourceColor, newColor) {
    if(sr < 0 || sc < 0 || sr >= image.length || sc >= image[sr].length || image[sr][sc] == newColor || image[sr][sc] !== sourceColor) {
       return;
    }
    image[sr][sc] = newColor;
    const directions = [[-1,0],[0,1],[1,0],[0,-1]];
    directions.forEach(d => {
        const [dsr, dsc] = d;
        const nsr = dsr + sr;
        const nsc = dsc + sc;
        fill(image, nsr, nsc, sourceColor, newColor); 
    });
    return image;
    
};