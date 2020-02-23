import { Num4 } from "./Num4";

export class Num5 extends Num4
{
    v: number = 0;

    constructor(x: number = 0, y: number = 0, z: number = 0, w: number = 0, v: number = 0)
    {
        super(x, y, z, w);

        this.v = v;
    }
}