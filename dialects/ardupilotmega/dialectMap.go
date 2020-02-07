package ardupilotmega

mavCmdMap := map[string]int{

	// Navigate to waypoint.
	"MAV_CMD_NAV_WAYPOINT":16,
	// Loiter around this waypoint an unlimited amount of time
	"MAV_CMD_NAV_LOITER_UNLIM":17,
	// Loiter around this waypoint for X turns
	"MAV_CMD_NAV_LOITER_TURNS":18,
	// Loiter around this waypoint for X seconds
	"MAV_CMD_NAV_LOITER_TIME":19,
	// Return to launch location
	"MAV_CMD_NAV_RETURN_TO_LAUNCH":20,
	// Land at location.
	"MAV_CMD_NAV_LAND":21,
	// Takeoff from ground / hand
	"MAV_CMD_NAV_TAKEOFF":22,
	// Land at local position (local frame only)
	"MAV_CMD_NAV_LAND_LOCAL":23,
	// Takeoff from local position (local frame only)
	"MAV_CMD_NAV_TAKEOFF_LOCAL":24,
	// Vehicle following, i.e. this waypoint represents the position of a moving vehicle
	"MAV_CMD_NAV_FOLLOW":25,
	// Continue on the current course and climb/descend to specified altitude.  When the altitude is reached continue to the next command (i.e., don't proceed to the next command until the desired altitude is reached.
	"MAV_CMD_NAV_CONTINUE_AND_CHANGE_ALT":30,
	// Begin loiter at the specified Latitude and Longitude.  If Lat=Lon=0, then loiter at the current position.  Don't consider the navigation command complete (don't leave loiter) until the altitude has been reached.  Additionally, if the Heading Required parameter is non-zero the  aircraft will not leave the loiter until heading toward the next waypoint.
	"MAV_CMD_NAV_LOITER_TO_ALT":31,
	// Begin following a target
	"MAV_CMD_DO_FOLLOW":32,
	// Reposition the "MAV after a follow target command has been sent
	"MAV_CMD_DO_FOLLOW_REPOSITION":33,
	// Start orbiting on the circumference of a circle defined by the parameters. Setting any value NaN results in using defaults.
	"MAV_CMD_DO_ORBIT":34,
	// Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	"MAV_CMD_NAV_ROI":80,
	// Control autonomous path planning on the MAV.
	"MAV_CMD_NAV_PATHPLANNING":81,
	// Navigate to waypoint using a spline path.
	"MAV_CMD_NAV_SPLINE_WAYPOINT":82,
	// Takeoff from ground using VTOL mode, and transition to forward flight with specified heading.
	"MAV_CMD_NAV_VTOL_TAKEOFF":84,
	// Land using VTOL mode
	"MAV_CMD_NAV_VTOL_LAND":85,
	// hand control over to an external controller
	"MAV_CMD_NAV_GUIDED_ENABLE":92,
	// Delay the next navigation command a number of seconds or until a specified time
	"MAV_CMD_NAV_DELAY":93,
	// Descend and place payload. Vehicle moves to specified location, descends until it detects a hanging payload has reached the ground, and then releases the payload. If ground is not detected before the reaching the maximum descent value (param1), the command will complete without releasing the payload.
	"MAV_CMD_NAV_PAYLOAD_PLACE":94,
	// NOP - This command is only used to mark the upper limit of the NAV/ACTION commands in the enumeration
	"MAV_CMD_NAV_LAST":95,
	// Delay mission state machine.
	"MAV_CMD_CONDITION_DELAY":112,
	// Ascend/descend at rate.  Delay mission state machine until desired altitude reached.
	"MAV_CMD_CONDITION_CHANGE_ALT":113,
	// Delay mission state machine until within desired distance of next NAV point.
	"MAV_CMD_CONDITION_DISTANCE":114,
	// Reach a certain target angle.
	"MAV_CMD_CONDITION_YAW":115,
	// NOP - This command is only used to mark the upper limit of the CONDITION commands in the enumeration
	"MAV_CMD_CONDITION_LAST":159,
	// Set system mode.
	"MAV_CMD_DO_SET_MODE":176,
	// Jump to the desired command in the mission list.  Repeat this action only the specified number of times
	"MAV_CMD_DO_JUMP":177,
	// Change speed and/or throttle set points.
	"MAV_CMD_DO_CHANGE_SPEED":178,
	// Changes the home location either to the current location or a specified location.
	"MAV_CMD_DO_SET_HOME":179,
	// Set a system parameter.  Caution!  Use of this command requires knowledge of the numeric enumeration value of the parameter.
	"MAV_CMD_DO_SET_PARAMETER":180,
	// Set a relay to a condition.
	"MAV_CMD_DO_SET_RELAY":181,
	// Cycle a relay on and off for a desired number of cycles with a desired period.
	"MAV_CMD_DO_REPEAT_RELAY":182,
	// Set a servo to a desired PWM value.
	"MAV_CMD_DO_SET_SERVO":183,
	// Cycle a between its nominal setting and a desired PWM for a desired number of cycles with a desired period.
	"MAV_CMD_DO_REPEAT_SERVO":184,
	// Terminate flight immediately
	"MAV_CMD_DO_FLIGHTTERMINATION":185,
	// Change altitude set point.
	"MAV_CMD_DO_CHANGE_ALTITUDE":186,
	// Mission command to perform a landing. This is used as a marker in a mission to tell the autopilot where a sequence of mission items that represents a landing starts. It may also be sent via a COMMAND_LONG to trigger a landing, in which case the nearest (geographically) landing sequence in the mission will be used. The Latitude/Longitude is optional, and may be set to 0 if not needed. If specified then it will be used to help find the closest landing sequence.
	"MAV_CMD_DO_LAND_START":189,
	// Mission command to perform a landing from a rally point.
	"MAV_CMD_DO_RALLY_LAND":190,
	// Mission command to safely abort an autonomous landing.
	"MAV_CMD_DO_GO_AROUND":191,
	// Reposition the vehicle to a specific WGS84 global position.
	"MAV_CMD_DO_REPOSITION":192,
	// If in a GPS controlled position mode, hold the current position or continue.
	"MAV_CMD_DO_PAUSE_CONTINUE":193,
	// Set moving direction to forward or reverse.
	"MAV_CMD_DO_SET_REVERSE":194,
	// Sets the region of interest (ROI) to a location. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	"MAV_CMD_DO_SET_ROI_LOCATION":195,
	// Sets the region of interest (ROI) to be toward next waypoint, with optional pitch/roll/yaw offset. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	"MAV_CMD_DO_SET_ROI_WPNEXT_OFFSET":196,
	// Cancels any previous ROI command returning the vehicle/sensors to default flight characteristics. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	"MAV_CMD_DO_SET_ROI_NONE":197,
	// Control onboard camera system.
	"MAV_CMD_DO_CONTROL_VIDEO":200,
	// Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	"MAV_CMD_DO_SET_ROI":201,
	// Configure digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ).
	"MAV_CMD_DO_DIGICAM_CONFIGURE":202,
	// Control digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ).
	"MAV_CMD_DO_DIGICAM_CONTROL":203,
	// Mission command to configure a camera or antenna mount
	"MAV_CMD_DO_MOUNT_CONFIGURE":204,
	// Mission command to control a camera or antenna mount
	"MAV_CMD_DO_MOUNT_CONTROL":205,
	// Mission command to set camera trigger distance for this flight. The camera is triggered each time this distance is exceeded. This command can also be used to set the shutter integration time for the camera.
	"MAV_CMD_DO_SET_CAM_TRIGG_DIST":206,
	// Mission command to enable the geofence
	"MAV_CMD_DO_FENCE_ENABLE":207,
	// Mission command to trigger a parachute
	"MAV_CMD_DO_PARACHUTE":208,
	// Mission command to perform motor test.
	"MAV_CMD_DO_MOTOR_TEST":209,
	// Change to/from inverted flight.
	"MAV_CMD_DO_INVERTED_FLIGHT":210,
	// Sets a desired vehicle turn angle and speed change.
	"MAV_CMD_NAV_SET_YAW_SPEED":213,
	// Mission command to set camera trigger interval for this flight. If triggering is enabled, the camera is triggered each time this interval expires. This command can also be used to set the shutter integration time for the camera.
	"MAV_CMD_DO_SET_CAM_TRIGG_INTERVAL":214,
	// Mission command to control a camera or antenna mount, using a quaternion as reference.
	"MAV_CMD_DO_MOUNT_CONTROL_QUAT":220,
	// set id of master controller
	"MAV_CMD_DO_GUIDED_MASTER":221,
	// Set limits for external control
	"MAV_CMD_DO_GUIDED_LIMITS":222,
	// Control vehicle engine. This is interpreted by the vehicles engine controller to change the target engine state. It is intended for vehicles with internal combustion engines
	"MAV_CMD_DO_ENGINE_CONTROL":223,
	// Set the mission item with sequence number seq as current item. This means that the "MAV will continue to this mission item on the shortest path (not following the mission items in-between).
	"MAV_CMD_DO_SET_MISSION_CURRENT":224,
	// NOP - This command is only used to mark the upper limit of the DO commands in the enumeration
	"MAV_CMD_DO_LAST":240,
	// Trigger calibration. This command will be only accepted if in pre-flight mode. Except for Temperature Calibration, only one sensor should be set in a single message and all others should be zero.
	"MAV_CMD_PREFLIGHT_CALIBRATION":241,
	// Set sensor offsets. This command will be only accepted if in pre-flight mode.
	"MAV_CMD_PREFLIGHT_SET_SENSOR_OFFSETS":242,
	// Trigger UAVCAN config. This command will be only accepted if in pre-flight mode.
	"MAV_CMD_PREFLIGHT_UAVCAN":243,
	// Request storage of different parameter values and logs. This command will be only accepted if in pre-flight mode.
	"MAV_CMD_PREFLIGHT_STORAGE":245,
	// Request the reboot or shutdown of system components.
	"MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN":246,
	// Override current mission with command to pause mission, pause mission and move to position, continue/resume mission. When param 1 indicates that the mission is paused ("MAV_GOTO_DO_HOLD), param 2 defines whether it holds in place or moves to another position.
	"MAV_CMD_OVERRIDE_GOTO":252,
	// start running a mission
	"MAV_CMD_MISSION_START":300,
	// Arms / Disarms a component
	"MAV_CMD_COMPONENT_ARM_DISARM":400,
	// Turns illuminators ON/OFF. An illuminator is a light source that is used for lighting up dark areas external to the sytstem: e.g. a torch or searchlight (as opposed to a light source for illuminating the system itself, e.g. an indicator light).
	"MAV_CMD_ILLUMINATOR_ON_OFF":405,
	// Request the home position from the vehicle.
	"MAV_CMD_GET_HOME_POSITION":410,
	// Starts receiver pairing.
	"MAV_CMD_START_RX_PAIR":500,
	// Request the interval between messages for a particular "MAVLink message ID. The receiver should ACK the command and then emit its response in a MESSAGE_INTERVAL message.
	"MAV_CMD_GET_MESSAGE_INTERVAL":510,
	// Set the interval between messages for a particular "MAVLink message ID. This interface replaces REQUEST_DATA_STREAM.
	"MAV_CMD_SET_MESSAGE_INTERVAL":511,
	// Request the target system(s) emit a single instance of a specified message (i.e. a "one-shot" version of_SET_MESSAGE_INTERVAL).
	"MAV_CMD_REQUEST_MESSAGE":512,
	// Request "MAVLink protocol version compatibility
	"MAV_CMD_REQUEST_PROTOCOL_VERSION":519,
	// Request autopilot capabilities. The receiver should ACK the command and then emit its capabilities in an AUTOPILOT_VERSION message
	"MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES":520,
	// Request camera information (CAMERA_INFORMATION).
	"MAV_CMD_REQUEST_CAMERA_INFORMATION":521,
	// Request camera settings (CAMERA_SETTINGS).
	"MAV_CMD_REQUEST_CAMERA_SETTINGS":522,
	// Request storage information (STORAGE_INFORMATION). Use the command's target_component to target a specific component's storage.
	"MAV_CMD_REQUEST_STORAGE_INFORMATION":525,
	// Format a storage medium. Once format is complete, a STORAGE_INFORMATION message is sent. Use the command's target_component to target a specific component's storage.
	"MAV_CMD_STORAGE_FORMAT":526,
	// Request camera capture status (CAMERA_CAPTURE_STATUS)
	"MAV_CMD_REQUEST_CAMERA_CAPTURE_STATUS":527,
	// Request flight information (FLIGHT_INFORMATION)
	"MAV_CMD_REQUEST_FLIGHT_INFORMATION":528,
	// Reset all camera settings to Factory Default
	"MAV_CMD_RESET_CAMERA_SETTINGS":529,
	// Set camera running mode. Use NaN for reserved values. GCS will send a_REQUEST_VIDEO_STREAM_STATUS command after a mode change if the camera supports video streaming.
	"MAV_CMD_SET_CAMERA_MODE":530,
	// Set camera zoom. Camera must respond with a CAMERA_SETTINGS message (on success). Use NaN for reserved values.
	"MAV_CMD_SET_CAMERA_ZOOM":531,
	// Set camera focus. Camera must respond with a CAMERA_SETTINGS message (on success). Use NaN for reserved values.
	"MAV_CMD_SET_CAMERA_FOCUS":532,
	// Tagged jump target. Can be jumped to with_DO_JUMP_TAG.
	"MAV_CMD_JUMP_TAG":600,
	// Jump to the matching tag in the mission list. Repeat this action for the specified number of times. A mission should contain a single matching tag for each jump. If this is not the case then a jump to a missing tag should complete the mission, and a jump where there are multiple matching tags should always select the one with the lowest mission sequence number.
	"MAV_CMD_DO_JUMP_TAG":601,
	// Start image capture sequence. Sends CAMERA_IMAGE_CAPTURED after each capture. Use NaN for reserved values.
	"MAV_CMD_IMAGE_START_CAPTURE":2000,
	// Stop image capture sequence Use NaN for reserved values.
	"MAV_CMD_IMAGE_STOP_CAPTURE":2001,
	// Re-request a CAMERA_IMAGE_CAPTURE message. Use NaN for reserved values.
	"MAV_CMD_REQUEST_CAMERA_IMAGE_CAPTURE":2002,
	// Enable or disable on-board camera triggering system.
	"MAV_CMD_DO_TRIGGER_CONTROL":2003,
	// Starts video capture (recording). Use NaN for reserved values.
	"MAV_CMD_VIDEO_START_CAPTURE":2500,
	// Stop the current video capture (recording). Use NaN for reserved values.
	"MAV_CMD_VIDEO_STOP_CAPTURE":2501,
	// Start video streaming
	"MAV_CMD_VIDEO_START_STREAMING":2502,
	// Stop the given video stream
	"MAV_CMD_VIDEO_STOP_STREAMING":2503,
	// Request video stream information (VIDEO_STREAM_INFORMATION)
	"MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION":2504,
	// Request video stream status (VIDEO_STREAM_STATUS)
	"MAV_CMD_REQUEST_VIDEO_STREAM_STATUS":2505,
	// Request to start streaming logging data over "MAVLink (see also LOGGING_DATA message)
	"MAV_CMD_LOGGING_START":2510,
	// Request to stop streaming log data over MAVLink
	"MAV_CMD_LOGGING_STOP":2511,
	//
	"MAV_CMD_AIRFRAME_CONFIGURATION":2520,
	// Request to start/stop transmitting over the high latency telemetry
	"MAV_CMD_CONTROL_HIGH_LATENCY":2600,
	// Create a panorama at the current position
	"MAV_CMD_PANORAMA_CREATE":2800,
	// Request VTOL transition
	"MAV_CMD_DO_VTOL_TRANSITION":3000,
	// Request authorization to arm the vehicle to a external entity, the arm authorizer is responsible to request all data that is needs from the vehicle before authorize or deny the request. If approved the progress of command_ack message should be set with period of time that this authorization is valid in seconds or in case it was denied it should be set with one of the reasons in ARM_AUTH_DENIED_REASON.
	"MAV_CMD_ARM_AUTHORIZATION_REQUEST":3001,
	// This command sets the submode to standard guided when vehicle is in guided mode. The vehicle holds position and altitude and the user can input the desired velocities along all three axes.
	"MAV_CMD_SET_GUIDED_SUBMODE_STANDARD":4000,
	// This command sets submode circle when vehicle is in guided mode. Vehicle flies along a circle facing the center of the circle. The user can input the velocity along the circle and change the radius. If no input is given the vehicle will hold position.
	"MAV_CMD_SET_GUIDED_SUBMODE_CIRCLE":4001,
	// Delay mission state machine until gate has been reached.
	"MAV_CMD_CONDITION_GATE":4501,
	// Fence return point. There can only be one fence return point.
	"MAV_CMD_NAV_FENCE_RETURN_POINT":5000,
	// Fence vertex for an inclusion polygon (the polygon must not be self-intersecting). The vehicle must stay within this area. Minimum of 3 vertices required.
	"MAV_CMD_NAV_FENCE_POLYGON_VERTEX_INCLUSION":5001,
	// Fence vertex for an exclusion polygon (the polygon must not be self-intersecting). The vehicle must stay outside this area. Minimum of 3 vertices required.
	"MAV_CMD_NAV_FENCE_POLYGON_VERTEX_EXCLUSION":5002,
	// Circular fence area. The vehicle must stay inside this area.
	"MAV_CMD_NAV_FENCE_CIRCLE_INCLUSION":5003,
	// Circular fence area. The vehicle must stay outside this area.
	"MAV_CMD_NAV_FENCE_CIRCLE_EXCLUSION":5004,
	// Rally point. You can have multiple rally points defined.
	"MAV_CMD_NAV_RALLY_POINT":5100,
	// Commands the vehicle to respond with a sequence of messages UAVCAN_NODE_INFO, one message per every UAVCAN node that is online. Note that some of the response messages can be lost, which the receiver can detect easily by checking whether every received UAVCAN_NODE_STATUS has a matching message UAVCAN_NODE_INFO received earlier; if not, this command should be sent again in order to request re-transmission of the node information messages.
	"MAV_CMD_UAVCAN_GET_NODE_INFO":5200,
	// Deploy payload on a Lat / Lon / Alt position. This includes the navigation to reach the required release position and velocity.
	"MAV_CMD_PAYLOAD_PREPARE_DEPLOY":30001,
	// Control the payload deployment.
	"MAV_CMD_PAYLOAD_CONTROL_DEPLOY":30002,
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	"MAV_CMD_WAYPOINT_USER_1":31000,
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	"MAV_CMD_WAYPOINT_USER_2":31001,
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	"MAV_CMD_WAYPOINT_USER_3":31002,
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	"MAV_CMD_WAYPOINT_USER_4":31003,
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	"MAV_CMD_WAYPOINT_USER_5":31004,
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	"MAV_CMD_SPATIAL_USER_1":31005,
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	"MAV_CMD_SPATIAL_USER_2":31006,
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	"MAV_CMD_SPATIAL_USER_3":31007,
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	"MAV_CMD_SPATIAL_USER_4":31008,
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	"MAV_CMD_SPATIAL_USER_5":31009,
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example:_DO_SET_PARAMETER item.
	"MAV_CMD_USER_1":31010,
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example:_DO_SET_PARAMETER item.
	"MAV_CMD_USER_2":31011,
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example:_DO_SET_PARAMETER item.
	"MAV_CMD_USER_3":31012,
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example:_DO_SET_PARAMETER item.
	"MAV_CMD_USER_4":31013,
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example:_DO_SET_PARAMETER item.
	"MAV_CMD_USER_5":31014,
	// Mission command to operate EPM gripper.
	"MAV_CMD_DO_GRIPPER":211,
	// Enable/disable autotune.
	"MAV_CMD_DO_AUTOTUNE_ENABLE":212,
	// Mission command to wait for an altitude or downwards vertical speed. This is meant for high altitude balloon launches, allowing the aircraft to be idle until either an altitude is reached or a negative vertical speed is reached (indicating early balloon burst). The wiggle time is how often to wiggle the control surfaces to prevent them seizing up.
	"MAV_CMD_NAV_ALTITUDE_WAIT":83,
	// A system wide power-off event has been initiated.
	"MAV_CMD_POWER_OFF_INITIATED":42000,
	// FLY button has been clicked.
	"MAV_CMD_SOLO_BTN_FLY_CLICK":42001,
	// FLY button has been held for 1.5 seconds.
	"MAV_CMD_SOLO_BTN_FLY_HOLD":42002,
	// PAUSE button has been clicked.
	"MAV_CMD_SOLO_BTN_PAUSE_CLICK":42003,
	// Magnetometer calibration based on fixed position        in earth field given by inclination, declination and intensity.
	"MAV_CMD_FIXED_MAG_CAL":42004,
	// Magnetometer calibration based on fixed expected field values in milliGauss.
	"MAV_CMD_FIXED_MAG_CAL_FIELD":42005,
	// Initiate a magnetometer calibration.
	"MAV_CMD_DO_START_MAG_CAL":42424,
	// Initiate a magnetometer calibration.
	"MAV_CMD_DO_ACCEPT_MAG_CAL":42425,
	// Cancel a running magnetometer calibration.
	"MAV_CMD_DO_CANCEL_MAG_CAL":42426,
	// Used when doing accelerometer calibration. When sent to the GCS tells it what position to put the vehicle in. When sent to the vehicle says what position the vehicle is in.
	"MAV_CMD_ACCELCAL_VEHICLE_POS":42429,
	// Reply with the version banner.
	"MAV_CMD_DO_SEND_BANNER":42428,
	// Command autopilot to get into factory test/diagnostic mode.
	"MAV_CMD_SET_FACTORY_TEST_MODE":42427,
	// Causes the gimbal to reset and boot as if it was just powered on.
	"MAV_CMD_GIMBAL_RESET":42501,
	// Reports progress and success or failure of gimbal axis calibration procedure.
	"MAV_CMD_GIMBAL_AXIS_CALIBRATION_STATUS":42502,
	// Starts commutation calibration on the gimbal.
	"MAV_CMD_GIMBAL_REQUEST_AXIS_CALIBRATION":42503,
	// Erases gimbal application and parameters.
	"MAV_CMD_GIMBAL_FULL_RESET":42505,
	// Command to operate winch.
	"MAV_CMD_DO_WINCH":42600,
	// Update the bootloader
	"MAV_CMD_FLASH_BOOTLOADER":42650,
	// Reset battery capacity for batteries that accumulate consumed battery via integration.
	"MAV_CMD_BATTERY_RESET":42651,
}
