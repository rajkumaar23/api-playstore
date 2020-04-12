<?php

/**
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */
class Utils
{
    public static function shouldUpdateCache($dateTime)
    {
        return empty($dateTime) || (time() > $dateTime + CACHE_INTERVAL_SECONDS);
    }
}
