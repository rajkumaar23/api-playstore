<?php
/**
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */

use Carbon\Carbon;

/**
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */
class Utils
{
    public static function shouldUpdateCache($dateTime)
    {
        return empty($dateTime) || (Carbon::parse($dateTime)->diffInMinutes(null) > CACHE_INTERVAL_MINUTES);
    }
}
