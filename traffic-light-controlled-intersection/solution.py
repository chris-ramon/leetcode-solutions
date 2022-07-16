from threading import Lock 

class TrafficLight:
    def __init__(self):
        self.current_green = 1
        self.mutex = Lock()

    def carArrived(
        self,
        carId: int,                      # ID of the car
        roadId: int,                     # ID of the road the car travels on. Can be 1 (road A) or 2 (road B)
        direction: int,                  # Direction of the car
        turnGreen: 'Callable[[], None]', # Use turnGreen() to turn light to green on current road
        crossCar: 'Callable[[], None]'   # Use crossCar() to make car cross the intersection
    ) -> None:
        self.mutex.acquire()
        if self.current_green == roadId:
            crossCar()
            self.mutex.release()
            return 

        turnGreen()
        crossCar()
        self.current_green = roadId
        self.mutex.release()

# TestCases
# [1,3,5,2,4]
# [2,1,2,4,3]
# [10,20,30,40,50]

# [1,2,3,4,5]
# [2,4,3,3,1]
# [10,20,30,40,40]
