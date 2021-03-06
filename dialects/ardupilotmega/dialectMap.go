package ardupilotmega

func FromString(cmdStr string) MAV_CMD {
	var cmd MAV_CMD

	switch cmdStr {
	// Navigate to waypoint.
	case "MAV_CMD_NAV_WAYPOINT":
		cmd = MAV_CMD_NAV_WAYPOINT
	// Loiter around this waypoint an unlimited amount of time
	case "MAV_CMD_NAV_LOITER_UNLIM":
		cmd = MAV_CMD_NAV_LOITER_UNLIM
	// Loiter around this waypoint for X turns
	case "MAV_CMD_NAV_LOITER_TURNS":
		cmd = MAV_CMD_NAV_LOITER_TURNS
	// Loiter around this waypoint for X seconds
	case "MAV_CMD_NAV_LOITER_TIME":
		cmd = MAV_CMD_NAV_LOITER_TIME
	// Return to launch location
	case "MAV_CMD_NAV_RETURN_TO_LAUNCH":
		cmd = MAV_CMD_NAV_RETURN_TO_LAUNCH
	// Land at location.
	case "MAV_CMD_NAV_LAND":
		cmd = MAV_CMD_NAV_LAND
	// Takeoff from ground / hand
	case "MAV_CMD_NAV_TAKEOFF":
		cmd = MAV_CMD_NAV_TAKEOFF
	// Land at local position (local frame only)
	case "MAV_CMD_NAV_LAND_LOCAL":
		cmd = MAV_CMD_NAV_LAND_LOCAL
	// Takeoff from local position (local frame only)
	case "MAV_CMD_NAV_TAKEOFF_LOCAL":
		cmd = MAV_CMD_NAV_TAKEOFF_LOCAL
	// Vehicle following, i.e. this waypoint represents the position of a moving vehicle
	case "MAV_CMD_NAV_FOLLOW":
		cmd = MAV_CMD_NAV_FOLLOW
	// Continue on the current course and climb/descend to specified altitude.  When the altitude is reached continue to the next command (i.e., don't proceed to the next command until the desired altitude is reached.
	case "MAV_CMD_NAV_CONTINUE_AND_CHANGE_ALT":
		cmd = MAV_CMD_NAV_CONTINUE_AND_CHANGE_ALT
	// Begin loiter at the specified Latitude and Longitude.  If Lat=Lon=0, then loiter at the current position.  Don't consider the navigation command complete (don't leave loiter) until the altitude has been reached.  Additionally, if the Heading Required parameter is non-zero the  aircraft will not leave the loiter until heading toward the next waypoint.
	case "MAV_CMD_NAV_LOITER_TO_ALT":
		cmd = MAV_CMD_NAV_LOITER_TO_ALT
	// Begin following a target
	case "MAV_CMD_DO_FOLLOW":
		cmd = MAV_CMD_DO_FOLLOW
	// Reposition the case "MAV after a follow target command has been sent
	case "MAV_CMD_DO_FOLLOW_REPOSITION":
		cmd = MAV_CMD_DO_FOLLOW_REPOSITION
	// Start orbiting on the circumference of a circle defined by the parameters. Setting any value NaN results in using defaults.
	case "MAV_CMD_DO_ORBIT":
		cmd = MAV_CMD_DO_ORBIT
	// Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	case "MAV_CMD_NAV_ROI":
		cmd = MAV_CMD_NAV_ROI
	// Control autonomous path planning on the MAV.
	case "MAV_CMD_NAV_PATHPLANNING":
		cmd = MAV_CMD_NAV_PATHPLANNING
	// Navigate to waypoint using a spline path.
	case "MAV_CMD_NAV_SPLINE_WAYPOINT":
		cmd = MAV_CMD_NAV_SPLINE_WAYPOINT
	// Takeoff from ground using VTOL mode, and transition to forward flight with specified heading.
	case "MAV_CMD_NAV_VTOL_TAKEOFF":
		cmd = MAV_CMD_NAV_VTOL_TAKEOFF
	// Land using VTOL mode
	case "MAV_CMD_NAV_VTOL_LAND":
		cmd = MAV_CMD_NAV_VTOL_LAND
	// hand control over to an external controller
	case "MAV_CMD_NAV_GUIDED_ENABLE":
		cmd = MAV_CMD_NAV_GUIDED_ENABLE
	// Delay the next navigation command a number of seconds or until a specified time
	case "MAV_CMD_NAV_DELAY":
		cmd = MAV_CMD_NAV_DELAY
	// Descend and place payload. Vehicle moves to specified location, descends until it detects a hanging payload has reached the ground, and then releases the payload. If ground is not detected before the reaching the maximum descent value (param1), the command will complete without releasing the payload.
	case "MAV_CMD_NAV_PAYLOAD_PLACE":
		cmd = MAV_CMD_NAV_PAYLOAD_PLACE
	// NOP - This command is only used to mark the upper limit of the NAV/ACTION commands in the enumeration
	case "MAV_CMD_NAV_LAST":
		cmd = MAV_CMD_NAV_LAST
	// Delay mission state machine.
	case "MAV_CMD_CONDITION_DELAY":
		cmd = MAV_CMD_CONDITION_DELAY
	// Ascend/descend at rate.  Delay mission state machine until desired altitude reached.
	case "MAV_CMD_CONDITION_CHANGE_ALT":
		cmd = MAV_CMD_CONDITION_CHANGE_ALT
	// Delay mission state machine until within desired distance of next NAV point.
	case "MAV_CMD_CONDITION_DISTANCE":
		cmd = MAV_CMD_CONDITION_DISTANCE
	// Reach a certain target angle.
	case "MAV_CMD_CONDITION_YAW":
		cmd = MAV_CMD_CONDITION_YAW
	// NOP - This command is only used to mark the upper limit of the CONDITION commands in the enumeration
	case "MAV_CMD_CONDITION_LAST":
		cmd = MAV_CMD_CONDITION_LAST
	// Set system mode.
	case "MAV_CMD_DO_SET_MODE":
		cmd = MAV_CMD_DO_SET_MODE
	// Jump to the desired command in the mission list.  Repeat this action only the specified number of times
	case "MAV_CMD_DO_JUMP":
		cmd = MAV_CMD_DO_JUMP
	// Change speed and/or throttle set points.
	case "MAV_CMD_DO_CHANGE_SPEED":
		cmd = MAV_CMD_DO_CHANGE_SPEED
	// Changes the home location either to the current location or a specified location.
	case "MAV_CMD_DO_SET_HOME":
		cmd = MAV_CMD_DO_SET_HOME
	// Set a system parameter.  Caution!  Use of this command requires knowledge of the numeric enumeration value of the parameter.
	case "MAV_CMD_DO_SET_PARAMETER":
		cmd = MAV_CMD_DO_SET_PARAMETER
	// Set a relay to a condition.
	case "MAV_CMD_DO_SET_RELAY":
		cmd = MAV_CMD_DO_SET_RELAY
	// Cycle a relay on and off for a desired number of cycles with a desired period.
	case "MAV_CMD_DO_REPEAT_RELAY":
		cmd = MAV_CMD_DO_REPEAT_RELAY
	// Set a servo to a desired PWM value.
	case "MAV_CMD_DO_SET_SERVO":
		cmd = MAV_CMD_DO_SET_SERVO
	// Cycle a between its nominal setting and a desired PWM for a desired number of cycles with a desired period.
	case "MAV_CMD_DO_REPEAT_SERVO":
		cmd = MAV_CMD_DO_REPEAT_SERVO
	// Terminate flight immediately
	case "MAV_CMD_DO_FLIGHTTERMINATION":
		cmd = MAV_CMD_DO_FLIGHTTERMINATION
	// Change altitude set point.
	case "MAV_CMD_DO_CHANGE_ALTITUDE":
		cmd = MAV_CMD_DO_CHANGE_ALTITUDE
	// Mission command to perform a landing. This is used as a marker in a mission to tell the autopilot where a sequence of mission items that represents a landing starts. It may also be sent via a COMMAND_LONG to trigger a landing, in which case the nearest (geographically) landing sequence in the mission will be used. The Latitude/Longitude is optional, and may be set to 0 if not needed. If specified then it will be used to help find the closest landing sequence.
	case "MAV_CMD_DO_LAND_START":
		cmd = MAV_CMD_DO_LAND_START
	// Mission command to perform a landing from a rally point.
	case "MAV_CMD_DO_RALLY_LAND":
		cmd = MAV_CMD_DO_RALLY_LAND
	// Mission command to safely abort an autonomous landing.
	case "MAV_CMD_DO_GO_AROUND":
		cmd = MAV_CMD_DO_GO_AROUND
	// Reposition the vehicle to a specific WGS84 global position.
	case "MAV_CMD_DO_REPOSITION":
		cmd = MAV_CMD_DO_REPOSITION
	// If in a GPS controlled position mode, hold the current position or continue.
	case "MAV_CMD_DO_PAUSE_CONTINUE":
		cmd = MAV_CMD_DO_PAUSE_CONTINUE
	// Set moving direction to forward or reverse.
	case "MAV_CMD_DO_SET_REVERSE":
		cmd = MAV_CMD_DO_SET_REVERSE
	// Sets the region of interest (ROI) to a location. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	case "MAV_CMD_DO_SET_ROI_LOCATION":
		cmd = MAV_CMD_DO_SET_ROI_LOCATION
	// Sets the region of interest (ROI) to be toward next waypoint, with optional pitch/roll/yaw offset. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	case "MAV_CMD_DO_SET_ROI_WPNEXT_OFFSET":
		cmd = MAV_CMD_DO_SET_ROI_WPNEXT_OFFSET
	// Cancels any previous ROI command returning the vehicle/sensors to default flight characteristics. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	case "MAV_CMD_DO_SET_ROI_NONE":
		cmd = MAV_CMD_DO_SET_ROI_NONE
	// Control onboard camera system.
	case "MAV_CMD_DO_CONTROL_VIDEO":
		cmd = MAV_CMD_DO_CONTROL_VIDEO
	// Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	case "MAV_CMD_DO_SET_ROI":
		cmd = MAV_CMD_DO_SET_ROI
	// Configure digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ).
	case "MAV_CMD_DO_DIGICAM_CONFIGURE":
		cmd = MAV_CMD_DO_DIGICAM_CONFIGURE
	// Control digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ).
	case "MAV_CMD_DO_DIGICAM_CONTROL":
		cmd = MAV_CMD_DO_DIGICAM_CONTROL
	// Mission command to configure a camera or antenna mount
	case "MAV_CMD_DO_MOUNT_CONFIGURE":
		cmd = MAV_CMD_DO_MOUNT_CONFIGURE
	// Mission command to control a camera or antenna mount
	case "MAV_CMD_DO_MOUNT_CONTROL":
		cmd = MAV_CMD_DO_MOUNT_CONTROL
	// Mission command to set camera trigger distance for this flight. The camera is triggered each time this distance is exceeded. This command can also be used to set the shutter integration time for the camera.
	case "MAV_CMD_DO_SET_CAM_TRIGG_DIST":
		cmd = MAV_CMD_DO_SET_CAM_TRIGG_DIST
	// Mission command to enable the geofence
	case "MAV_CMD_DO_FENCE_ENABLE":
		cmd = MAV_CMD_DO_FENCE_ENABLE
	// Mission command to trigger a parachute
	case "MAV_CMD_DO_PARACHUTE":
		cmd = MAV_CMD_DO_PARACHUTE
	// Mission command to perform motor test.
	case "MAV_CMD_DO_MOTOR_TEST":
		cmd = MAV_CMD_DO_MOTOR_TEST
	// Change to/from inverted flight.
	case "MAV_CMD_DO_INVERTED_FLIGHT":
		cmd = MAV_CMD_DO_INVERTED_FLIGHT
	// Sets a desired vehicle turn angle and speed change.
	case "MAV_CMD_NAV_SET_YAW_SPEED":
		cmd = MAV_CMD_NAV_SET_YAW_SPEED
	// Mission command to set camera trigger interval for this flight. If triggering is enabled, the camera is triggered each time this interval expires. This command can also be used to set the shutter integration time for the camera.
	case "MAV_CMD_DO_SET_CAM_TRIGG_INTERVAL":
		cmd = MAV_CMD_DO_SET_CAM_TRIGG_INTERVAL
	// Mission command to control a camera or antenna mount, using a quaternion as reference.
	case "MAV_CMD_DO_MOUNT_CONTROL_QUAT":
		cmd = MAV_CMD_DO_MOUNT_CONTROL_QUAT
	// set id of master controller
	case "MAV_CMD_DO_GUIDED_MASTER":
		cmd = MAV_CMD_DO_GUIDED_MASTER
	// Set limits for external control
	case "MAV_CMD_DO_GUIDED_LIMITS":
		cmd = MAV_CMD_DO_GUIDED_LIMITS
	// Control vehicle engine. This is interpreted by the vehicles engine controller to change the target engine state. It is intended for vehicles with internal combustion engines
	case "MAV_CMD_DO_ENGINE_CONTROL":
		cmd = MAV_CMD_DO_ENGINE_CONTROL
	// Set the mission item with sequence number seq as current item. This means that the case "MAV will continue to this mission item on the shortest path (not following the mission items in-between).
	case "MAV_CMD_DO_SET_MISSION_CURRENT":
		cmd = MAV_CMD_DO_SET_MISSION_CURRENT
	// NOP - This command is only used to mark the upper limit of the DO commands in the enumeration
	case "MAV_CMD_DO_LAST":
		cmd = MAV_CMD_DO_LAST
	// Trigger calibration. This command will be only accepted if in pre-flight mode. Except for Temperature Calibration, only one sensor should be set in a single message and all others should be zero.
	case "MAV_CMD_PREFLIGHT_CALIBRATION":
		cmd = MAV_CMD_PREFLIGHT_CALIBRATION
	// Set sensor offsets. This command will be only accepted if in pre-flight mode.
	case "MAV_CMD_PREFLIGHT_SET_SENSOR_OFFSETS":
		cmd = MAV_CMD_PREFLIGHT_SET_SENSOR_OFFSETS
	// Trigger UAVCAN config. This command will be only accepted if in pre-flight mode.
	case "MAV_CMD_PREFLIGHT_UAVCAN":
		cmd = MAV_CMD_PREFLIGHT_UAVCAN
	// Request storage of different parameter values and logs. This command will be only accepted if in pre-flight mode.
	case "MAV_CMD_PREFLIGHT_STORAGE":
		cmd = MAV_CMD_PREFLIGHT_STORAGE
	// Request the reboot or shutdown of system components.
	case "MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN":
		cmd = MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN
	// Override current mission with command to pause mission, pause mission and move to position, continue/resume mission. When param 1 indicates that the mission is paused (case "MAV_GOTO_DO_HOLD), param 2 defines whether it holds in place or moves to another position.
	case "MAV_CMD_OVERRIDE_GOTO":
		cmd = MAV_CMD_OVERRIDE_GOTO
	// start running a mission
	case "MAV_CMD_MISSION_START":
		cmd = MAV_CMD_MISSION_START
	// Arms / Disarms a component
	case "MAV_CMD_COMPONENT_ARM_DISARM":
		cmd = MAV_CMD_COMPONENT_ARM_DISARM
	// Turns illuminators ON/OFF. An illuminator is a light source that is used for lighting up dark areas external to the sytstem: e.g. a torch or searchlight (as opposed to a light source for illuminating the system itself, e.g. an indicator light).
	case "MAV_CMD_ILLUMINATOR_ON_OFF":
		cmd = MAV_CMD_ILLUMINATOR_ON_OFF
	// Request the home position from the vehicle.
	case "MAV_CMD_GET_HOME_POSITION":
		cmd = MAV_CMD_GET_HOME_POSITION
	// Starts receiver pairing.
	case "MAV_CMD_START_RX_PAIR":
		cmd = MAV_CMD_START_RX_PAIR
	// Request the interval between messages for a particular case "MAVLink message ID. The receiver should ACK the command and then emit its response in a MESSAGE_INTERVAL message.
	case "MAV_CMD_GET_MESSAGE_INTERVAL":
		cmd = MAV_CMD_GET_MESSAGE_INTERVAL
	// Set the interval between messages for a particular case "MAVLink message ID. This interface replaces REQUEST_DATA_STREAM.
	case "MAV_CMD_SET_MESSAGE_INTERVAL":
		cmd = MAV_CMD_SET_MESSAGE_INTERVAL
	// Request the target system(s) emit a single instance of a specified message (i.e. a "one-shot" version of_SET_MESSAGE_INTERVAL).
	case "MAV_CMD_REQUEST_MESSAGE":
		cmd = MAV_CMD_REQUEST_MESSAGE
	// Request case "MAVLink protocol version compatibility
	case "MAV_CMD_REQUEST_PROTOCOL_VERSION":
		cmd = MAV_CMD_REQUEST_PROTOCOL_VERSION
	// Request autopilot capabilities. The receiver should ACK the command and then emit its capabilities in an AUTOPILOT_VERSION message
	case "MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES":
		cmd = MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES
	// Request camera information (CAMERA_INFORMATION).
	case "MAV_CMD_REQUEST_CAMERA_INFORMATION":
		cmd = MAV_CMD_REQUEST_CAMERA_INFORMATION
	// Request camera settings (CAMERA_SETTINGS).
	case "MAV_CMD_REQUEST_CAMERA_SETTINGS":
		cmd = MAV_CMD_REQUEST_CAMERA_SETTINGS
	// Request storage information (STORAGE_INFORMATION). Use the command's target_component to target a specific component's storage.
	case "MAV_CMD_REQUEST_STORAGE_INFORMATION":
		cmd = MAV_CMD_REQUEST_STORAGE_INFORMATION
	// Format a storage medium. Once format is complete, a STORAGE_INFORMATION message is sent. Use the command's target_component to target a specific component's storage.
	case "MAV_CMD_STORAGE_FORMAT":
		cmd = MAV_CMD_STORAGE_FORMAT
	// Request camera capture status (CAMERA_CAPTURE_STATUS)
	case "MAV_CMD_REQUEST_CAMERA_CAPTURE_STATUS":
		cmd = MAV_CMD_REQUEST_CAMERA_CAPTURE_STATUS
	// Request flight information (FLIGHT_INFORMATION)
	case "MAV_CMD_REQUEST_FLIGHT_INFORMATION":
		cmd = MAV_CMD_REQUEST_FLIGHT_INFORMATION
	// Reset all camera settings to Factory Default
	case "MAV_CMD_RESET_CAMERA_SETTINGS":
		cmd = MAV_CMD_RESET_CAMERA_SETTINGS
	// Set camera running mode. Use NaN for reserved values. GCS will send a_REQUEST_VIDEO_STREAM_STATUS command after a mode change if the camera supports video streaming.
	case "MAV_CMD_SET_CAMERA_MODE":
		cmd = MAV_CMD_SET_CAMERA_MODE
	// Set camera zoom. Camera must respond with a CAMERA_SETTINGS message (on success). Use NaN for reserved values.
	case "MAV_CMD_SET_CAMERA_ZOOM":
		cmd = MAV_CMD_SET_CAMERA_ZOOM
	// Set camera focus. Camera must respond with a CAMERA_SETTINGS message (on success). Use NaN for reserved values.
	case "MAV_CMD_SET_CAMERA_FOCUS":
		cmd = MAV_CMD_SET_CAMERA_FOCUS
	// Tagged jump target. Can be jumped to with_DO_JUMP_TAG.
	case "MAV_CMD_JUMP_TAG":
		cmd = MAV_CMD_JUMP_TAG
	// Jump to the matching tag in the mission list. Repeat this action for the specified number of times. A mission should contain a single matching tag for each jump. If this is not the case then a jump to a missing tag should complete the mission, and a jump where there are multiple matching tags should always select the one with the lowest mission sequence number.
	case "MAV_CMD_DO_JUMP_TAG":
		cmd = MAV_CMD_DO_JUMP_TAG
	// Start image capture sequence. Sends CAMERA_IMAGE_CAPTURED after each capture. Use NaN for reserved values.
	case "MAV_CMD_IMAGE_START_CAPTURE":
		cmd = MAV_CMD_IMAGE_START_CAPTURE
	// Stop image capture sequence Use NaN for reserved values.
	case "MAV_CMD_IMAGE_STOP_CAPTURE":
		cmd = MAV_CMD_IMAGE_STOP_CAPTURE
	// Re-request a CAMERA_IMAGE_CAPTURE message. Use NaN for reserved values.
	case "MAV_CMD_REQUEST_CAMERA_IMAGE_CAPTURE":
		cmd = MAV_CMD_REQUEST_CAMERA_IMAGE_CAPTURE
	// Enable or disable on-board camera triggering system.
	case "MAV_CMD_DO_TRIGGER_CONTROL":
		cmd = MAV_CMD_DO_TRIGGER_CONTROL
	// Starts video capture (recording). Use NaN for reserved values.
	case "MAV_CMD_VIDEO_START_CAPTURE":
		cmd = MAV_CMD_VIDEO_START_CAPTURE
	// Stop the current video capture (recording). Use NaN for reserved values.
	case "MAV_CMD_VIDEO_STOP_CAPTURE":
		cmd = MAV_CMD_VIDEO_STOP_CAPTURE
	// Start video streaming
	case "MAV_CMD_VIDEO_START_STREAMING":
		cmd = MAV_CMD_VIDEO_START_STREAMING
	// Stop the given video stream
	case "MAV_CMD_VIDEO_STOP_STREAMING":
		cmd = MAV_CMD_VIDEO_STOP_STREAMING
	// Request video stream information (VIDEO_STREAM_INFORMATION)
	case "MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION":
		cmd = MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION
	// Request video stream status (VIDEO_STREAM_STATUS)
	case "MAV_CMD_REQUEST_VIDEO_STREAM_STATUS":
		cmd = MAV_CMD_REQUEST_VIDEO_STREAM_STATUS
	// Request to start streaming logging data over case "MAVLink (see also LOGGING_DATA message)
	case "MAV_CMD_LOGGING_START":
		cmd = MAV_CMD_LOGGING_START
	// Request to stop streaming log data over MAVLink
	case "MAV_CMD_LOGGING_STOP":
		cmd = MAV_CMD_LOGGING_STOP
	//
	case "MAV_CMD_AIRFRAME_CONFIGURATION":
		cmd = MAV_CMD_AIRFRAME_CONFIGURATION
	// Request to start/stop transmitting over the high latency telemetry
	case "MAV_CMD_CONTROL_HIGH_LATENCY":
		cmd = MAV_CMD_CONTROL_HIGH_LATENCY
	// Create a panorama at the current position
	case "MAV_CMD_PANORAMA_CREATE":
		cmd = MAV_CMD_PANORAMA_CREATE
	// Request VTOL transition
	case "MAV_CMD_DO_VTOL_TRANSITION":
		cmd = MAV_CMD_DO_VTOL_TRANSITION
	// Request authorization to arm the vehicle to a external entity, the arm authorizer is responsible to request all data that is needs from the vehicle before authorize or deny the request. If approved the progress of command_ack message should be set with period of time that this authorization is valid in seconds or in case it was denied it should be set with one of the reasons in ARM_AUTH_DENIED_REASON.
	case "MAV_CMD_ARM_AUTHORIZATION_REQUEST":
		cmd = MAV_CMD_ARM_AUTHORIZATION_REQUEST
	// This command sets the submode to standard guided when vehicle is in guided mode. The vehicle holds position and altitude and the user can input the desired velocities along all three axes.
	case "MAV_CMD_SET_GUIDED_SUBMODE_STANDARD":
		cmd = MAV_CMD_SET_GUIDED_SUBMODE_STANDARD
	// This command sets submode circle when vehicle is in guided mode. Vehicle flies along a circle facing the center of the circle. The user can input the velocity along the circle and change the radius. If no input is given the vehicle will hold position.
	case "MAV_CMD_SET_GUIDED_SUBMODE_CIRCLE":
		cmd = MAV_CMD_SET_GUIDED_SUBMODE_CIRCLE
	// Delay mission state machine until gate has been reached.
	case "MAV_CMD_CONDITION_GATE":
		cmd = MAV_CMD_CONDITION_GATE
	// Fence return point. There can only be one fence return point.
	case "MAV_CMD_NAV_FENCE_RETURN_POINT":
		cmd = MAV_CMD_NAV_FENCE_RETURN_POINT
	// Fence vertex for an inclusion polygon (the polygon must not be self-intersecting). The vehicle must stay within this area. Minimum of 3 vertices required.
	case "MAV_CMD_NAV_FENCE_POLYGON_VERTEX_INCLUSION":
		cmd = MAV_CMD_NAV_FENCE_POLYGON_VERTEX_INCLUSION
	// Fence vertex for an exclusion polygon (the polygon must not be self-intersecting). The vehicle must stay outside this area. Minimum of 3 vertices required.
	case "MAV_CMD_NAV_FENCE_POLYGON_VERTEX_EXCLUSION":
		cmd = MAV_CMD_NAV_FENCE_POLYGON_VERTEX_EXCLUSION
	// Circular fence area. The vehicle must stay inside this area.
	case "MAV_CMD_NAV_FENCE_CIRCLE_INCLUSION":
		cmd = MAV_CMD_NAV_FENCE_CIRCLE_INCLUSION
	// Circular fence area. The vehicle must stay outside this area.
	case "MAV_CMD_NAV_FENCE_CIRCLE_EXCLUSION":
		cmd = MAV_CMD_NAV_FENCE_CIRCLE_EXCLUSION
	// Rally point. You can have multiple rally points defined.
	case "MAV_CMD_NAV_RALLY_POINT":
		cmd = MAV_CMD_NAV_RALLY_POINT
	// Commands the vehicle to respond with a sequence of messages UAVCAN_NODE_INFO, one message per every UAVCAN node that is online. Note that some of the response messages can be lost, which the receiver can detect easily by checking whether every received UAVCAN_NODE_STATUS has a matching message UAVCAN_NODE_INFO received earlier; if not, this command should be sent again in order to request re-transmission of the node information messages.
	case "MAV_CMD_UAVCAN_GET_NODE_INFO":
		cmd = MAV_CMD_UAVCAN_GET_NODE_INFO
	// Deploy payload on a Lat / Lon / Alt position. This includes the navigation to reach the required release position and velocity.
	case "MAV_CMD_PAYLOAD_PREPARE_DEPLOY":
		cmd = MAV_CMD_PAYLOAD_PREPARE_DEPLOY
	// Control the payload deployment.
	case "MAV_CMD_PAYLOAD_CONTROL_DEPLOY":
		cmd = MAV_CMD_PAYLOAD_CONTROL_DEPLOY
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	case "MAV_CMD_WAYPOINT_USER_1":
		cmd = MAV_CMD_WAYPOINT_USER_1
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	case "MAV_CMD_WAYPOINT_USER_2":
		cmd = MAV_CMD_WAYPOINT_USER_2
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	case "MAV_CMD_WAYPOINT_USER_3":
		cmd = MAV_CMD_WAYPOINT_USER_3
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	case "MAV_CMD_WAYPOINT_USER_4":
		cmd = MAV_CMD_WAYPOINT_USER_4
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	case "MAV_CMD_WAYPOINT_USER_5":
		cmd = MAV_CMD_WAYPOINT_USER_5
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	case "MAV_CMD_SPATIAL_USER_1":
		cmd = MAV_CMD_SPATIAL_USER_1
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	case "MAV_CMD_SPATIAL_USER_2":
		cmd = MAV_CMD_SPATIAL_USER_2
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	case "MAV_CMD_SPATIAL_USER_3":
		cmd = MAV_CMD_SPATIAL_USER_3
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	case "MAV_CMD_SPATIAL_USER_4":
		cmd = MAV_CMD_SPATIAL_USER_4
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	case "MAV_CMD_SPATIAL_USER_5":
		cmd = MAV_CMD_SPATIAL_USER_5
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example:_DO_SET_PARAMETER item.
	case "MAV_CMD_USER_1":
		cmd = MAV_CMD_USER_1
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example:_DO_SET_PARAMETER item.
	case "MAV_CMD_USER_2":
		cmd = MAV_CMD_USER_2
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example:_DO_SET_PARAMETER item.
	case "MAV_CMD_USER_3":
		cmd = MAV_CMD_USER_3
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example:_DO_SET_PARAMETER item.
	case "MAV_CMD_USER_4":
		cmd = MAV_CMD_USER_4
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example:_DO_SET_PARAMETER item.
	case "MAV_CMD_USER_5":
		cmd = MAV_CMD_USER_5
	// Mission command to operate EPM gripper.
	case "MAV_CMD_DO_GRIPPER":
		cmd = MAV_CMD_DO_GRIPPER
	// Enable/disable autotune.
	case "MAV_CMD_DO_AUTOTUNE_ENABLE":
		cmd = MAV_CMD_DO_AUTOTUNE_ENABLE
	// Mission command to wait for an altitude or downwards vertical speed. This is meant for high altitude balloon launches, allowing the aircraft to be idle until either an altitude is reached or a negative vertical speed is reached (indicating early balloon burst). The wiggle time is how often to wiggle the control surfaces to prevent them seizing up.
	case "MAV_CMD_NAV_ALTITUDE_WAIT":
		cmd = MAV_CMD_NAV_ALTITUDE_WAIT
	// A system wide power-off event has been initiated.
	case "MAV_CMD_POWER_OFF_INITIATED":
		cmd = MAV_CMD_POWER_OFF_INITIATED
	// FLY button has been clicked.
	case "MAV_CMD_SOLO_BTN_FLY_CLICK":
		cmd = MAV_CMD_SOLO_BTN_FLY_CLICK
	// FLY button has been held for 1.5 seconds.
	case "MAV_CMD_SOLO_BTN_FLY_HOLD":
		cmd = MAV_CMD_SOLO_BTN_FLY_HOLD
	// PAUSE button has been clicked.
	case "MAV_CMD_SOLO_BTN_PAUSE_CLICK":
		cmd = MAV_CMD_SOLO_BTN_PAUSE_CLICK
	// Magnetometer calibration based on fixed position        in earth field given by inclination, declination and intensity.
	case "MAV_CMD_FIXED_MAG_CAL":
		cmd = MAV_CMD_FIXED_MAG_CAL
	// Magnetometer calibration based on fixed expected field values in milliGauss.
	case "MAV_CMD_FIXED_MAG_CAL_FIELD":
		cmd = MAV_CMD_FIXED_MAG_CAL_FIELD
	// Initiate a magnetometer calibration.
	case "MAV_CMD_DO_START_MAG_CAL":
		cmd = MAV_CMD_DO_START_MAG_CAL
	// Initiate a magnetometer calibration.
	case "MAV_CMD_DO_ACCEPT_MAG_CAL":
		cmd = MAV_CMD_DO_ACCEPT_MAG_CAL
	// Cancel a running magnetometer calibration.
	case "MAV_CMD_DO_CANCEL_MAG_CAL":
		cmd = MAV_CMD_DO_CANCEL_MAG_CAL
	// Used when doing accelerometer calibration. When sent to the GCS tells it what position to put the vehicle in. When sent to the vehicle says what position the vehicle is in.
	case "MAV_CMD_ACCELCAL_VEHICLE_POS":
		cmd = MAV_CMD_ACCELCAL_VEHICLE_POS
	// Reply with the version banner.
	case "MAV_CMD_DO_SEND_BANNER":
		cmd = MAV_CMD_DO_SEND_BANNER
	// Command autopilot to get into factory test/diagnostic mode.
	case "MAV_CMD_SET_FACTORY_TEST_MODE":
		cmd = MAV_CMD_SET_FACTORY_TEST_MODE
	// Causes the gimbal to reset and boot as if it was just powered on.
	case "MAV_CMD_GIMBAL_RESET":
		cmd = MAV_CMD_GIMBAL_RESET
	// Reports progress and success or failure of gimbal axis calibration procedure.
	case "MAV_CMD_GIMBAL_AXIS_CALIBRATION_STATUS":
		cmd = MAV_CMD_GIMBAL_AXIS_CALIBRATION_STATUS
	// Starts commutation calibration on the gimbal.
	case "MAV_CMD_GIMBAL_REQUEST_AXIS_CALIBRATION":
		cmd = MAV_CMD_GIMBAL_REQUEST_AXIS_CALIBRATION
	// Erases gimbal application and parameters.
	case "MAV_CMD_GIMBAL_FULL_RESET":
		cmd = MAV_CMD_GIMBAL_FULL_RESET
	// Command to operate winch.
	case "MAV_CMD_DO_WINCH":
		cmd = MAV_CMD_DO_WINCH
	// Update the bootloader
	case "MAV_CMD_FLASH_BOOTLOADER":
		cmd = MAV_CMD_FLASH_BOOTLOADER
	// Reset battery capacity for batteries that accumulate consumed battery via integration.
	case "MAV_CMD_BATTERY_RESET":
		cmd = MAV_CMD_BATTERY_RESET
	default:
		cmd = MAV_CMD_NAV_WAYPOINT
	}

	return cmd
}
