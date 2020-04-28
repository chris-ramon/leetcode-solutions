/**
 * Definition for Employee.
 * function Employee(id, importance, subordinates) {
 *     this.id = id;
 *     this.importance = importance;
 *     this.subordinates = subordinates;
 * }
 */

/**
 * @param {Employee[]} employees
 * @param {number} id
 * @return {number}
 */
var GetImportance = function(employees, id) {
    if(!employees || !employees.length || !id) return 0;
    
    var totalImportance = 0;
    var employeesMap = new Map();
    employees.forEach(e => employeesMap.set(e.id, e) );
    
    var queue = [id];
    while(queue.length) {
        var size = queue.length;
        for(var i = 0; i < size; i++) {
            var id = queue.shift();
            if(!employeesMap.has(id)) continue;
            var employee = employeesMap.get(id);
            if(employee.subordinates.length) {
                queue.push(...employee.subordinates);
            }
            totalImportance += employee.importance;
        }
    }
    
    return totalImportance;
};