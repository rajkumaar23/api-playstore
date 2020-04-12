<?php
/**
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */

class DB
{
    private $conn;

    public function __construct()
    {

        $client = new MongoDB\Client(getenv("PLAYSTORE_MONGO_URI"));
        $this->conn = $client->db->data;
    }

    /**
     * @return mixed
     */
    public function getConn()
    {
        return $this->conn;
    }
}
